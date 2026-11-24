package api_runner

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"api-tests-template/internal/utils"
)

const (
	maxAttempts    = 5
	retryBaseDelay = 300 * time.Millisecond
	retryLogLabel  = "Retry"
)

// retryTransport повторяет запрос, если стенд ответил 503. Так nginx стенда отвечает
// на всплеск запросов, например при очистке десятков объявлений подряд
type retryTransport struct {
	next http.RoundTripper
}

func (rt retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// apitest копирует запрос без GetBody, поэтому тело читаем заранее и отправляем заново при повторе
	var body []byte
	if req.Body != nil {
		var err error
		body, err = io.ReadAll(req.Body)
		_ = req.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	for attempt := 1; ; attempt++ {
		attemptReq := req.Clone(req.Context())
		attemptReq.Body = http.NoBody
		if len(body) > 0 {
			attemptReq.Body = io.NopCloser(bytes.NewReader(body))
		}

		resp, err := rt.next.RoundTrip(attemptReq)
		if err != nil || resp.StatusCode != http.StatusServiceUnavailable || attempt == maxAttempts {
			return resp, err
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()

		delay := retryBaseDelay * time.Duration(attempt)
		utils.LogWithLabelAndTimeStamp(retryLogLabel, fmt.Sprintf("%s %s: HTTP 503, повтор через %v (попытка %d из %d)",
			req.Method, req.URL.Path, delay, attempt+1, maxAttempts))

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(delay):
		}
	}
}

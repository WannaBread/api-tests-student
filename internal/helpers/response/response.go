// Package response содержит общие хелперы для проверки и разбора HTTP-ответов.
package response

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// RequireStatus читает тело ответа и проверяет статус-код; при несовпадении тело попадает в сообщение об ошибке
func RequireStatus(t *testing.T, resp *http.Response, expectedStatus int, action string) []byte {
	t.Helper()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoErrorf(t, err, "%s: не удалось прочитать тело ответа", action)

	require.Equalf(t, expectedStatus, resp.StatusCode,
		"%s: ожидали HTTP %d, получили HTTP %d. Тело ответа: %s",
		action, expectedStatus, resp.StatusCode, string(body))
	return body
}

// DecodeJSON проверяет статус-код ответа и разбирает JSON-тело в структуру T
func DecodeJSON[T any](t *testing.T, resp *http.Response, expectedStatus int, action string) T {
	t.Helper()
	body := RequireStatus(t, resp, expectedStatus, action)

	var result T
	require.NoErrorf(t, json.Unmarshal(body, &result),
		"%s: не удалось разобрать JSON ответа: %s", action, string(body))
	return result
}

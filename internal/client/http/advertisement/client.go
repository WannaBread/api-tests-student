package advertisement

import (
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/advertisement/models"
)

const (
	formFieldTitle       = "title"
	formFieldDescription = "description"
	formFieldPrice       = "price"
	formFieldQuantity    = "quantity"
	formFieldPhotos      = "photos"
	queryParamID         = "id"
)

// HttpPostAdvertisement выполняет multipart POST-запрос создания объявления; незаданные поля не передаются
func HttpPostAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) *http.Response {
	t.Helper()
	httpRequest := apiRunner.GetRunner().Auth(token).Create().
		Post(path.AdvertisementPath)

	if request.Title != "" {
		httpRequest = httpRequest.MultipartFormData(formFieldTitle, request.Title)
	}
	if request.Description != "" {
		httpRequest = httpRequest.MultipartFormData(formFieldDescription, request.Description)
	}
	if request.Price != nil {
		httpRequest = httpRequest.MultipartFormData(formFieldPrice, strconv.FormatFloat(*request.Price, 'f', -1, 64))
	}
	if request.Quantity != nil {
		httpRequest = httpRequest.MultipartFormData(formFieldQuantity, strconv.Itoa(*request.Quantity))
	}
	if len(request.Photos) > 0 {
		httpRequest = httpRequest.MultipartFile(formFieldPhotos, request.Photos...)
	}

	return httpRequest.Expect(t).End().Response
}

// HttpGetAdvertisement выполняет GET-запрос получения объявления по id (ручка публичная)
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().
		Get(path.AdvertisementPath).
		Query(queryParamID, id).
		Expect(t).
		End().Response
}

// HttpDeleteAdvertisement выполняет DELETE-запрос удаления объявления по id
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(path.AdvertisementPath).
		Query(queryParamID, id).
		Expect(t).
		End().Response
}

package photos

import (
	"net/http"
	"testing"

	photosClient "api-tests-template/internal/client/http/photos"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/advertisement/models"
)

// GetAdvertisementPhotos возвращает список фотографий объявления и проверяет, что вернулся статус 200
func GetAdvertisementPhotos(t *testing.T, token, advertisementID string) []models.PhotoResponse {
	t.Helper()
	rawResponse := photosClient.HttpGetAdvertisementPhotos(t, token, advertisementID)
	return response.DecodeJSON[[]models.PhotoResponse](t, rawResponse, http.StatusOK, "Получение фотографий объявления")
}

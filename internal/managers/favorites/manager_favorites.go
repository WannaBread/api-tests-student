package favorites

import (
	"net/http"
	"testing"

	favoritesClient "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/response"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/managers/favorites/models"
)

const (
	actionGet         = "Получение списка избранного"
	actionGetInvalid  = "Получение списка избранного с ошибкой"
	actionAdd         = "Добавление объявления в избранное"
	actionAddInvalid  = "Добавление объявления в избранное с ошибкой"
	actionBulkAdd     = "Пакетное добавление объявлений в избранное"
	actionRemove      = "Удаление объявления из избранного"
	actionRemoveError = "Удаление объявления из избранного с ошибкой"
)

// GetFavorites возвращает избранное пользователя одной страницей максимального размера (вмещает весь лимит избранного)
func GetFavorites(t *testing.T, token string) models.FavoriteListResponse {
	t.Helper()
	rawResponse := favoritesClient.HttpGetFavorites(t, token, limits.MaxPageSize)
	return response.DecodeJSON[models.FavoriteListResponse](t, rawResponse, http.StatusOK, actionGet)
}

// GetFavoritesExpectingError запрашивает избранное и возвращает тело ошибки с ожидаемым статусом
func GetFavoritesExpectingError(t *testing.T, token string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	rawResponse := favoritesClient.HttpGetFavorites(t, token, limits.MaxPageSize)
	return response.DecodeJSON[commonModels.ErrorResponse](t, rawResponse, expectedStatus, actionGetInvalid)
}

// AddFavorite добавляет объявление в избранное и проверяет, что вернулся статус 201
func AddFavorite(t *testing.T, token, advertisementID string) models.FavoriteItem {
	t.Helper()
	rawResponse := favoritesClient.HttpPostFavorite(t, token, models.AddFavoriteRequest{AdvertisementID: advertisementID})
	return response.DecodeJSON[models.FavoriteItem](t, rawResponse, http.StatusCreated, actionAdd)
}

// AddFavoriteExpectingError пытается добавить объявление в избранное и возвращает тело ошибки с ожидаемым статусом
func AddFavoriteExpectingError(t *testing.T, token, advertisementID string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	rawResponse := favoritesClient.HttpPostFavorite(t, token, models.AddFavoriteRequest{AdvertisementID: advertisementID})
	return response.DecodeJSON[commonModels.ErrorResponse](t, rawResponse, expectedStatus, actionAddInvalid)
}

// BulkAddFavorites добавляет пачку объявлений в избранное и проверяет, что вернулся статус 200
func BulkAddFavorites(t *testing.T, token string, advertisementIDs []string) models.BulkAddFavoritesResult {
	t.Helper()
	rawResponse := favoritesClient.HttpPostFavoritesBulk(t, token,
		models.BulkAddFavoritesRequest{AdvertisementIDs: advertisementIDs})
	return response.DecodeJSON[models.BulkAddFavoritesResult](t, rawResponse, http.StatusOK, actionBulkAdd)
}

// RemoveFavorite удаляет объявление из избранного и проверяет, что вернулся статус 204
func RemoveFavorite(t *testing.T, token, advertisementID string) {
	t.Helper()
	rawResponse := favoritesClient.HttpDeleteFavorite(t, token, advertisementID)
	response.RequireStatus(t, rawResponse, http.StatusNoContent, actionRemove)
}

// RemoveFavoriteExpectingError пытается удалить объявление из избранного и возвращает тело ошибки с ожидаемым статусом
func RemoveFavoriteExpectingError(t *testing.T, token, advertisementID string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	rawResponse := favoritesClient.HttpDeleteFavorite(t, token, advertisementID)
	return response.DecodeJSON[commonModels.ErrorResponse](t, rawResponse, expectedStatus, actionRemoveError)
}

package advertisements

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/advertisements/models"
)

// GetAdvertisementsList возвращает типизированный результат поиска.
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	t.Helper()
	raw := advertisements.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()

	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать ответ поиска")

	var searchResult models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &searchResult), "Не удалось разобрать ответ поиска")
	return searchResult
}

// GetAdvertisementsPageOldestFirst возвращает страницу каталога с сортировкой от старых к новым.
// Для первой страницы cursor пустой.
func GetAdvertisementsPageOldestFirst(t *testing.T, token string, limit int, cursor string) models.AdvertisementsList {
	t.Helper()
	rawResponse := advertisements.HttpGetAdvertisementsPageOldestFirst(t, token, limit, cursor)
	return response.DecodeJSON[models.AdvertisementsList](t, rawResponse, http.StatusOK, "Получение страницы каталога объявлений")
}

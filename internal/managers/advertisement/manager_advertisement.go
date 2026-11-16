package advertisement

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	advertisementClient "api-tests-template/internal/client/http/advertisement"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/advertisement/models"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/utils"
)

const (
	// TitlePrefix помечает объявления, созданные автотестами
	TitlePrefix         = "autotest "
	randomTitleLength   = 12
	DefaultDescription  = "Объявление создано автотестом"
	DefaultPrice        = 2499.99
	DefaultQuantity     = 5
	DefaultPhotosCount  = 1
	cleanupLogLabel     = "Cleanup"
	actionCreate        = "Создание объявления"
	actionGet           = "Получение объявления по id"
	actionDelete        = "Удаление объявления"
	actionCreateInvalid = "Создание объявления с ошибкой"
	actionGetInvalid    = "Получение объявления по id с ошибкой"
)

// NewAdvertisementRequest собирает валидный запрос на создание объявления со всеми параметрами и уникальным названием
func NewAdvertisementRequest(photosCount int) models.CreateAdvertisementRequest {
	price := DefaultPrice
	quantity := DefaultQuantity
	return models.CreateAdvertisementRequest{
		Title:       TitlePrefix + utils.RandomString(randomTitleLength),
		Description: DefaultDescription,
		Price:       &price,
		Quantity:    &quantity,
		Photos:      fixtures.PhotoPaths(photosCount),
	}
}

// CreateAdvertisement создаёт объявление и проверяет, что вернулся статус 201
func CreateAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) models.AdvertisementResponse {
	t.Helper()
	rawResponse := advertisementClient.HttpPostAdvertisement(t, token, request)
	return response.DecodeJSON[models.AdvertisementResponse](t, rawResponse, http.StatusCreated, actionCreate)
}

// CreateTemporaryAdvertisement создаёт объявление и регистрирует его удаление после завершения теста
func CreateTemporaryAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) models.AdvertisementResponse {
	t.Helper()
	created := CreateAdvertisement(t, token, request)
	t.Cleanup(func() {
		DeleteAdvertisementIfExists(t, token, created.ID)
	})
	return created
}

// CreateAdvertisementExpectingError отправляет запрос на создание объявления и возвращает тело ошибки с ожидаемым статусом.
// Если API всё же создало объявление, оно сразу удаляется, чтобы упавший тест не оставлял данных на стенде.
func CreateAdvertisementExpectingError(t *testing.T, token string, request models.CreateAdvertisementRequest,
	expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	rawResponse := advertisementClient.HttpPostAdvertisement(t, token, request)
	if rawResponse.StatusCode == http.StatusCreated {
		created := response.DecodeJSON[models.AdvertisementResponse](t, rawResponse, http.StatusCreated, actionCreate)
		DeleteAdvertisementIfExists(t, token, created.ID)
		require.Failf(t, "Объявление создано, хотя запрос невалидный",
			"%s: ожидали HTTP %d, получили HTTP %d. Созданное объявление %s удалено",
			actionCreateInvalid, expectedStatus, http.StatusCreated, created.ID)
	}
	return response.DecodeJSON[commonModels.ErrorResponse](t, rawResponse, expectedStatus, actionCreateInvalid)
}

// GetAdvertisement получает объявление по id и проверяет, что вернулся статус 200
func GetAdvertisement(t *testing.T, id string) models.AdvertisementResponse {
	t.Helper()
	rawResponse := advertisementClient.HttpGetAdvertisement(t, id)
	return response.DecodeJSON[models.AdvertisementResponse](t, rawResponse, http.StatusOK, actionGet)
}

// GetAdvertisementExpectingError запрашивает объявление по id и возвращает тело ошибки с ожидаемым статусом
func GetAdvertisementExpectingError(t *testing.T, id string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	rawResponse := advertisementClient.HttpGetAdvertisement(t, id)
	return response.DecodeJSON[commonModels.ErrorResponse](t, rawResponse, expectedStatus, actionGetInvalid)
}

// DeleteAdvertisement удаляет объявление и проверяет, что вернулся статус 204
func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	rawResponse := advertisementClient.HttpDeleteAdvertisement(t, token, id)
	response.RequireStatus(t, rawResponse, http.StatusNoContent, actionDelete)
}

// DeleteAdvertisementIfExists удаляет объявление при очистке данных: не валит тест, если объявление уже удалено,
// а неудачную попытку удаления только логирует
func DeleteAdvertisementIfExists(t *testing.T, token, id string) {
	t.Helper()
	getResponse := advertisementClient.HttpGetAdvertisement(t, id)
	_ = getResponse.Body.Close()
	if getResponse.StatusCode == http.StatusNotFound {
		utils.LogWithLabelAndTimeStamp(cleanupLogLabel, fmt.Sprintf("Объявление %s уже удалено", id))
		return
	}

	deleteResponse := advertisementClient.HttpDeleteAdvertisement(t, token, id)
	_ = deleteResponse.Body.Close()
	if deleteResponse.StatusCode != http.StatusNoContent {
		utils.LogWithLabelAndTimeStamp(cleanupLogLabel,
			fmt.Sprintf("Не удалось удалить объявление %s: HTTP %d", id, deleteResponse.StatusCode))
		return
	}
	utils.LogWithLabelAndTimeStamp(cleanupLogLabel, fmt.Sprintf("Объявление %s удалено", id))
}

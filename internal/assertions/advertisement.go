package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement/models"
)

// AdvertisementMatchesRequest проверяет, что объявление содержит параметры, переданные при создании
func AdvertisementMatchesRequest(t *testing.T, actual models.AdvertisementResponse, expected models.CreateAdvertisementRequest) {
	t.Helper()
	require.NotEmpty(t, actual.ID, "У объявления пустой id")
	require.Equal(t, expected.Title, actual.Title, "title не совпадает с переданным при создании")
	require.Equal(t, expected.Description, actual.Description, "description не совпадает с переданным при создании")
	if expected.Price != nil {
		require.Equal(t, *expected.Price, actual.Price, "price не совпадает с переданным при создании")
	}
	if expected.Quantity != nil {
		require.Equal(t, *expected.Quantity, actual.Quantity, "quantity не совпадает с переданным при создании")
	}
	require.Lenf(t, actual.Photos, len(expected.Photos),
		"Количество фотографий у объявления не совпадает с количеством загруженных (%d)", len(expected.Photos))
}

// AdvertisementEqualsCreated проверяет, что полученное объявление совпадает с созданным: id, владелец и фотографии
func AdvertisementEqualsCreated(t *testing.T, actual, created models.AdvertisementResponse) {
	t.Helper()
	require.Equal(t, created.ID, actual.ID, "id объявления не совпадает с созданным")
	require.Equal(t, created.UserID, actual.UserID, "user_id объявления не совпадает с создателем")
	require.Equal(t, created.CreatedAt, actual.CreatedAt, "created_at не совпадает с ответом на создание")
	PhotosEqual(t, actual.Photos, created.Photos)
}

// PhotosEqual проверяет, что списки фотографий совпадают по id, url и порядку сортировки
func PhotosEqual(t *testing.T, actual, expected []models.PhotoResponse) {
	t.Helper()
	require.Lenf(t, actual, len(expected), "Ожидали %d фотографий, получили %d", len(expected), len(actual))

	actualByID := make(map[string]models.PhotoResponse, len(actual))
	for _, photo := range actual {
		actualByID[photo.ID] = photo
	}
	for _, expectedPhoto := range expected {
		actualPhoto, found := actualByID[expectedPhoto.ID]
		require.Truef(t, found, "Фотография с id %s не найдена", expectedPhoto.ID)
		require.NotEmptyf(t, actualPhoto.URL, "У фотографии %s пустой url", expectedPhoto.ID)
		require.Equalf(t, expectedPhoto.URL, actualPhoto.URL, "url фотографии %s не совпадает", expectedPhoto.ID)
		require.Equalf(t, expectedPhoto.SortOrder, actualPhoto.SortOrder,
			"sort_order фотографии %s не совпадает", expectedPhoto.ID)
	}
}

package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/favorites/models"
)

// FavoritesContainAdvertisement проверяет, что объявление есть в списке избранного
func FavoritesContainAdvertisement(t *testing.T, favorites models.FavoriteListResponse, advertisementID string) {
	t.Helper()
	require.Truef(t, favoritesHaveAdvertisement(favorites, advertisementID),
		"Объявления %s нет в избранном, а должно быть. Всего в избранном: %d", advertisementID, favorites.Total)
}

// FavoritesDoNotContainAdvertisement проверяет, что объявления нет в списке избранного
func FavoritesDoNotContainAdvertisement(t *testing.T, favorites models.FavoriteListResponse, advertisementID string) {
	t.Helper()
	require.Falsef(t, favoritesHaveAdvertisement(favorites, advertisementID),
		"Объявление %s осталось в избранном, хотя должно было пропасть", advertisementID)
}

// favoritesHaveAdvertisement ищет объявление в списке избранного по id
func favoritesHaveAdvertisement(favorites models.FavoriteListResponse, advertisementID string) bool {
	for _, item := range favorites.Items {
		if item.Advertisement.ID == advertisementID {
			return true
		}
	}
	return false
}

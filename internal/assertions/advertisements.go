package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

// ContainsAdvertisementByTitle проверяет наличие объявления с точным названием.
func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()
	for _, item := range items {
		if item.Title == title {
			return
		}
	}
	require.Failf(t, "Объявление не найдено", "В результатах поиска нет объявления с названием %q", title)
}

package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

// ContainsAdvertisementByID проверяет наличие объявления с указанным id в результатах поиска.
func ContainsAdvertisementByID(t *testing.T, items []models.Advertisement, id string) {
	t.Helper()
	for _, item := range items {
		if item.ID == id {
			return
		}
	}
	require.Failf(t, "Объявление не найдено", "В результатах поиска (%d шт.) нет объявления с id %s", len(items), id)
}

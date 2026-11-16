package models

import advertisementModels "api-tests-template/internal/managers/advertisement/models"

// FavoriteItem — запись избранного; возвращается в POST /favorites и в списке GET /favorites
type FavoriteItem struct {
	FavoriteID    string                                    `json:"favorite_id"`
	AddedAt       string                                    `json:"added_at"`
	Advertisement advertisementModels.AdvertisementResponse `json:"advertisement"`
}

// FavoriteListResponse — страница избранного пользователя
type FavoriteListResponse struct {
	Items      []FavoriteItem `json:"items"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	Total      int            `json:"total"`
	TotalPages int            `json:"total_pages"`
}

// BulkAddFavoritesResult — результат POST /favorites/bulk
type BulkAddFavoritesResult struct {
	Added        []string `json:"added"`
	Skipped      []string `json:"skipped"`
	NotFound     []string `json:"not_found"`
	LimitReached bool     `json:"limit_reached"`
}

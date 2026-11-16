package models

// AddFavoriteRequest — тело запроса POST /favorites
type AddFavoriteRequest struct {
	AdvertisementID string `json:"advertisement_id"`
}

// BulkAddFavoritesRequest — тело запроса POST /favorites/bulk
type BulkAddFavoritesRequest struct {
	AdvertisementIDs []string `json:"advertisement_ids"`
}

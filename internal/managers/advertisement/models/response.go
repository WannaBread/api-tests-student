package models

// AdvertisementResponse — объявление в ответах POST и GET /advertisement
type AdvertisementResponse struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Quantity    int             `json:"quantity"`
	UserID      string          `json:"user_id"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Photos      []PhotoResponse `json:"photos"`
}

// PhotoResponse — фотография объявления
type PhotoResponse struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	SortOrder int    `json:"sort_order"`
	CreatedAt string `json:"created_at"`
}

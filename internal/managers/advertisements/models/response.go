package models

// AdvertisementsList — страница результатов поиска объявлений.
type AdvertisementsList struct {
	Items      []Advertisement `json:"items"`
	HasNext    bool            `json:"has_next"`
	NextCursor string          `json:"next_cursor"`
	Total      int             `json:"total"`
}

// Advertisement содержит поля объявления, необходимые для проверки поиска.
type Advertisement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

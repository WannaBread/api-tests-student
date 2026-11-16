package models

// CreateAdvertisementRequest — поля multipart-формы создания объявления.
// Незаданные поля (пустая строка, nil, пустой список) не передаются в запросе.
type CreateAdvertisementRequest struct {
	Title       string
	Description string
	Price       *float64
	Quantity    *int
	// Photos — пути к файлам фотографий на диске
	Photos []string
}

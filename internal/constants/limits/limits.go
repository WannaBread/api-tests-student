// Package limits содержит бизнес-ограничения API TestBoard, проверенные на стенде.
package limits

const (
	MaxTitleLength       = 50
	MaxDescriptionLength = 500
	MinPrice             = 0
	MaxPrice             = 1000000
	MinQuantity          = 1
	MaxQuantity          = 100
	MaxPhotosCount       = 3
	MaxFavoritesPerUser  = 50
	// MaxPageSize — максимальный размер страницы каталога и списка избранного
	MaxPageSize = 50
)

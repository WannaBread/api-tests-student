// Package apierrors содержит коды и тексты ошибок, которые возвращает API TestBoard.
package apierrors

// Значения поля error в ответе об ошибке
const (
	CodeUnauthorized   = "unauthorized"
	CodeInvalidRequest = "invalid_request"
	CodeNotFound       = "not_found"
	CodeForbidden      = "forbidden"
	CodeAlreadyExists  = "already_exists"
	CodeLimitReached   = "limit_reached"
)

// Тексты ошибок авторизации
const (
	MessageAuthHeaderRequired = "Authorization header required"
	MessageInvalidToken       = "Invalid or expired token" //nolint:gosec // текст ошибки API, а не секрет
)

// Тексты ошибок валидации при создании объявления
const (
	MessageTitleRequired = "title is required"
	// MessageTitleTooLongTemplate — параметры: максимальная длина, переданная длина
	MessageTitleTooLongTemplate     = "title must not exceed %d characters provided %d"
	MessageDescriptionRequired      = "description is required"
	MessageDescriptionLengthInvalid = "description must be between 1 and 500 characters"
	MessagePriceOutOfRange          = "price must be between 0 and 1000000 inclusive"
	MessageQuantityOutOfRange       = "quantity must be between 1 and 100 inclusive"
	MessagePhotoRequired            = "at least 1 photo is required"
	MessageTooManyPhotos            = "maximum 3 photos allowed"
)

// Тексты ошибок ручек избранного
const (
	MessageAdvertisementNotFound    = "Advertisement not found"
	MessageFavoriteNotFound         = "Favorite entry not found"
	MessageAlreadyInFavorites       = "Advertisement is already in favorites"
	MessageOwnAdvertisementFavorite = "Cannot add own advertisement to favorites"
	MessageFavoritesLimitReached    = "Favorites limit of 50 reached"
	MessageInvalidAdvertisementID   = "Invalid advertisement ID"
	// MessageInvalidUUIDLengthTemplate — параметр: длина переданного идентификатора
	MessageInvalidUUIDLengthTemplate = "invalid UUID length: %d"
)

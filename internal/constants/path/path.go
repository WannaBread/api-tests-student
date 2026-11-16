package path

const (
	AuthLoginPath        = "/auth/login"
	AuthRegisterPath     = "/auth/register"
	MyAdvertisementsPath = "/my/advertisements"
	AdvertisementsPath   = "/advertisements"
	AdvertisementPath    = "/advertisement"
	FavoritesPath        = "/favorites"
	FavoritesBulkPath    = "/favorites/bulk"

	// AdvertisementPhotosPathTemplate — шаблон пути, параметр: id объявления
	AdvertisementPhotosPathTemplate = "/advertisements/%s/photos"
	// FavoriteByAdIDPathTemplate — шаблон пути, параметр: id объявления
	FavoriteByAdIDPathTemplate = "/favorites/%s"
)

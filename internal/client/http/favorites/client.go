package favorites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/favorites/models"
)

// HttpGetFavorites выполняет GET-запрос первой страницы избранного указанного размера
func HttpGetFavorites(t *testing.T, token string, pageSize int) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.FavoritesPath).
		Query("page_size", strconv.Itoa(pageSize)).
		Expect(t).
		End().Response
}

// HttpPostFavorite выполняет POST-запрос добавления объявления в избранное
func HttpPostFavorite(t *testing.T, token string, request models.AddFavoriteRequest) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Post(path.FavoritesPath).
		ContentType("application/json").
		JSON(request).
		Expect(t).
		End().Response
}

// HttpPostFavoritesBulk выполняет POST-запрос пакетного добавления объявлений в избранное
func HttpPostFavoritesBulk(t *testing.T, token string, request models.BulkAddFavoritesRequest) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Post(path.FavoritesBulkPath).
		ContentType("application/json").
		JSON(request).
		Expect(t).
		End().Response
}

// HttpDeleteFavorite выполняет DELETE-запрос удаления объявления из избранного
func HttpDeleteFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(fmt.Sprintf(path.FavoriteByAdIDPathTemplate, url.PathEscape(advertisementID))).
		Expect(t).
		End().Response
}

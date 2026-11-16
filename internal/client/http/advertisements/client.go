package advertisements

import (
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpGetAdvertisements ищет объявления по подстроке в названии.
func HttpGetAdvertisements(t *testing.T, token, textForSearch string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.AdvertisementsPath).
		Query("search", textForSearch).
		Expect(t).
		Status(http.StatusOK).
		End().Response
}

// HttpGetAdvertisementsPageOldestFirst запрашивает страницу общего каталога, отсортированного от старых объявлений к новым;
// пустой cursor — первая страница.
func HttpGetAdvertisementsPageOldestFirst(t *testing.T, token string, limit int, cursor string) *http.Response {
	t.Helper()
	request := apiRunner.GetRunner().Auth(token).Create().
		Get(path.AdvertisementsPath).
		Query("limit", strconv.Itoa(limit)).
		Query("sort_by", "created_at").
		Query("sort_order", "asc")
	if cursor != "" {
		request = request.Query("cursor", cursor)
	}
	return request.Expect(t).
		Status(http.StatusOK).
		End().Response
}

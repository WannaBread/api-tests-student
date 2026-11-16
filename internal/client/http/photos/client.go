package photos

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpGetAdvertisementPhotos выполняет GET-запрос списка фотографий объявления
func HttpGetAdvertisementPhotos(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Get(fmt.Sprintf(path.AdvertisementPhotosPathTemplate, url.PathEscape(advertisementID))).
		Expect(t).
		End().Response
}

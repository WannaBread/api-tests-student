package advertisementCreate

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/advertisements"
	photosManager "api-tests-template/internal/managers/photos"
	base "api-tests-template/tests"
)

// Positive: создание объявления со всеми параметрами и максимальным количеством фотографий
func (s *TestSuite) TestCreateAdvertisementWithAllParameters() {
	request := advertisement.NewAdvertisementRequest(limits.MaxPhotosCount)
	base.Step(fmt.Sprintf("Создаём объявление %q: description=%q, price=%v, quantity=%d, фотографий: %d",
		request.Title, request.Description, *request.Price, *request.Quantity, len(request.Photos)))
	created := advertisement.CreateTemporaryAdvertisement(s.T(), s.owner.Token, request)
	base.Check(fmt.Sprintf("Объявление создано, id=%s", created.ID))

	s.Run("Ответ на создание содержит переданные параметры", func() {
		assertions.AdvertisementMatchesRequest(s.T(), created, request)
		require.Equal(s.T(), s.owner.User.Id, created.UserID, "Объявление создано не от имени авторизованного пользователя")
		base.Check("Параметры в ответе на создание совпадают с переданными")
	})

	s.Run("GET /advertisement возвращает те же параметры, что передали при создании", func() {
		base.Step(fmt.Sprintf("Запрашиваем объявление по id=%s", created.ID))
		fetched := advertisement.GetAdvertisement(s.T(), created.ID)
		assertions.AdvertisementMatchesRequest(s.T(), fetched, request)
		assertions.AdvertisementEqualsCreated(s.T(), fetched, created)
		base.Check("Объявление по id совпадает с созданным")
	})

	s.Run("Объявление находится поиском по полному названию", func() {
		base.Step(fmt.Sprintf("Ищем объявления по названию %q", request.Title))
		list := advertisements.GetAdvertisementsList(s.T(), s.owner.Token, request.Title)
		assertions.ContainsAdvertisementByID(s.T(), list.Items, created.ID)
		base.Check("Созданное объявление есть в результатах поиска")
	})

	s.Run("Фотографии загружены и доступны через GET /advertisements/{id}/photos", func() {
		base.Step(fmt.Sprintf("Запрашиваем фотографии объявления id=%s", created.ID))
		photos := photosManager.GetAdvertisementPhotos(s.T(), s.owner.Token, created.ID)
		require.Lenf(s.T(), photos, len(request.Photos), "Ожидали %d загруженных фотографий", len(request.Photos))
		assertions.PhotosEqual(s.T(), photos, created.Photos)
		base.Check(fmt.Sprintf("Все %d фотографии доступны и совпадают с загруженными", len(photos)))
	})
}

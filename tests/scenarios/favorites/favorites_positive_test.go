package favorites

import (
	"net/http"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// Positive: объявление появляется в избранном после добавления и пропадает после удаления из избранного
func (s *TestSuite) TestAddAndRemoveFavorite() {
	ad := s.createOwnerAdvertisement()

	addedToFavorites := s.Run("Добавленное объявление появляется в списке избранного", func() {
		base.Step("Покупатель добавляет объявление в избранное")
		favorite := favorites.AddFavorite(s.T(), s.buyer.Token, ad.ID)
		require.Equal(s.T(), ad.ID, favorite.Advertisement.ID, "В ответе на добавление вернулось другое объявление")
		require.NotEmpty(s.T(), favorite.FavoriteID, "В ответе на добавление пустой favorite_id")

		base.Step("Запрашиваем список избранного покупателя")
		assertions.FavoritesContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.ID)
		base.Check("Объявление есть в избранном")
	})
	require.True(s.T(), addedToFavorites, "Не удалось добавить объявление в избранное, удаление проверить нельзя")

	s.Run("Удалённое из избранного объявление пропадает из списка", func() {
		base.Step("Покупатель удаляет объявление из избранного")
		favorites.RemoveFavorite(s.T(), s.buyer.Token, ad.ID)

		base.Step("Запрашиваем список избранного покупателя")
		assertions.FavoritesDoNotContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.ID)
		base.Check("Объявления больше нет в избранном")
	})
}

// Positive: объявление, удалённое владельцем через DELETE /advertisement, пропадает из избранного
func (s *TestSuite) TestDeletedAdvertisementDisappearsFromFavorites() {
	ad := s.createOwnerAdvertisement()

	base.Precondition("Покупатель добавляет объявление в избранное")
	favorites.AddFavorite(s.T(), s.buyer.Token, ad.ID)
	assertions.FavoritesContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.ID)

	base.Step("Владелец удаляет объявление")
	advertisement.DeleteAdvertisement(s.T(), s.owner.Token, ad.ID)

	s.Run("Объявление удалено", func() {
		errorResponse := advertisement.GetAdvertisementExpectingError(s.T(), ad.ID, http.StatusNotFound)
		assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeNotFound, apierrors.MessageAdvertisementNotFound)
		base.Check("GET /advertisement возвращает 404")
	})

	s.Run("Удалённое объявление пропало из избранного", func() {
		base.Step("Запрашиваем список избранного покупателя")
		assertions.FavoritesDoNotContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.ID)
		base.Check("Объявления больше нет в избранном")
	})
}

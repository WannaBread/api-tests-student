package favorites

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// Negative: повторное добавление того же объявления в избранное возвращает 409
func (s *TestSuite) TestAddSameAdvertisementTwice() {
	ad := s.createOwnerAdvertisement()

	base.Precondition("Покупатель добавляет объявление в избранное первый раз")
	favorites.AddFavorite(s.T(), s.buyer.Token, ad.ID)

	base.Step("Покупатель добавляет то же объявление повторно")
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token, ad.ID, http.StatusConflict)

	assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeAlreadyExists, apierrors.MessageAlreadyInFavorites)
	base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusConflict, errorResponse.Message))
}

// Negative: добавление в избранное собственного объявления возвращает 403
func (s *TestSuite) TestAddOwnAdvertisement() {
	ad := s.createOwnerAdvertisement()

	base.Step("Владелец добавляет в избранное своё объявление")
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.owner.Token, ad.ID, http.StatusForbidden)
	assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeForbidden, apierrors.MessageOwnAdvertisementFavorite)
	base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusForbidden, errorResponse.Message))

	s.Run("Объявление не попало в избранное владельца", func() {
		assertions.FavoritesDoNotContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.owner.Token), ad.ID)
		base.Check("Объявления нет в избранном владельца")
	})
}

// Negative: добавление в избранное сверх лимита MaxFavoritesPerUser возвращает 400
func (s *TestSuite) TestAddFavoriteOverLimit() {
	base.Precondition("Регистрируем отдельного пользователя, чтобы заполнение лимита не влияло на другие тесты")
	limitedUser := auth.RegisterRandomUser(s.T())

	base.Precondition(fmt.Sprintf("Заполняем избранное пользователя до лимита %d объявлениями из каталога",
		limits.MaxFavoritesPerUser))
	favorites.FillFavoritesToLimit(s.T(), limitedUser.Token)

	// Объявление создаётся после заполнения лимита, чтобы не попасть в избранное вместе с объявлениями каталога
	ad := s.createOwnerAdvertisement()

	base.Step("Пользователь добавляет в избранное ещё одно объявление")
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), limitedUser.Token, ad.ID, http.StatusBadRequest)
	assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeLimitReached, apierrors.MessageFavoritesLimitReached)
	base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusBadRequest, errorResponse.Message))

	s.Run("Объявление сверх лимита не попало в избранное", func() {
		favoritesList := favorites.GetFavorites(s.T(), limitedUser.Token)
		require.Equal(s.T(), limits.MaxFavoritesPerUser, favoritesList.Total, "Количество избранного превысило лимит")
		assertions.FavoritesDoNotContainAdvertisement(s.T(), favoritesList, ad.ID)
		base.Check(fmt.Sprintf("В избранном ровно %d объявлений, нового среди них нет", limits.MaxFavoritesPerUser))
	})
}

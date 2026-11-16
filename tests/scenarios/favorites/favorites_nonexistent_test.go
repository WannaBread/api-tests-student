package favorites

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

// invalidAdvertisementID — идентификатор не в формате UUID
const invalidAdvertisementID = "not-a-uuid"

// nonexistentIDCase описывает операцию с несуществующим или некорректным ad_id и ожидаемую ошибку
type nonexistentIDCase struct {
	name            string
	advertisementID string
	expectedStatus  int
	expectedCode    string
	expectedMessage string
}

// Negative: добавление в избранное несуществующего или некорректного ad_id
func (s *TestSuite) TestAddFavoriteWithNonexistentAdvertisement() {
	cases := []nonexistentIDCase{
		{
			name:            "Несуществующий ad_id",
			advertisementID: utils.RandomUUID(),
			expectedStatus:  http.StatusNotFound,
			expectedCode:    apierrors.CodeNotFound,
			expectedMessage: apierrors.MessageAdvertisementNotFound,
		},
		{
			name:            "ad_id не в формате UUID",
			advertisementID: invalidAdvertisementID,
			expectedStatus:  http.StatusBadRequest,
			expectedCode:    apierrors.CodeInvalidRequest,
			expectedMessage: fmt.Sprintf(apierrors.MessageInvalidUUIDLengthTemplate, len(invalidAdvertisementID)),
		},
		{
			name:            "Пустой ad_id",
			advertisementID: "",
			expectedStatus:  http.StatusBadRequest,
			expectedCode:    apierrors.CodeInvalidRequest,
			expectedMessage: fmt.Sprintf(apierrors.MessageInvalidUUIDLengthTemplate, 0),
		},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Step(fmt.Sprintf("Добавляем в избранное ad_id=%q", tc.advertisementID))
			errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token, tc.advertisementID, tc.expectedStatus)

			assertions.ErrorResponseEquals(s.T(), errorResponse, tc.expectedCode, tc.expectedMessage)
			base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", tc.expectedStatus, errorResponse.Message))
		})
	}
}

// Negative: удаление из избранного несуществующего или некорректного ad_id
func (s *TestSuite) TestRemoveFavoriteWithNonexistentAdvertisement() {
	cases := []nonexistentIDCase{
		{
			name:            "Несуществующий ad_id",
			advertisementID: utils.RandomUUID(),
			expectedStatus:  http.StatusNotFound,
			expectedCode:    apierrors.CodeNotFound,
			expectedMessage: apierrors.MessageFavoriteNotFound,
		},
		{
			name:            "ad_id не в формате UUID",
			advertisementID: invalidAdvertisementID,
			expectedStatus:  http.StatusBadRequest,
			expectedCode:    apierrors.CodeInvalidRequest,
			expectedMessage: apierrors.MessageInvalidAdvertisementID,
		},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Step(fmt.Sprintf("Удаляем из избранного ad_id=%q", tc.advertisementID))
			errorResponse := favorites.RemoveFavoriteExpectingError(s.T(), s.buyer.Token, tc.advertisementID, tc.expectedStatus)

			assertions.ErrorResponseEquals(s.T(), errorResponse, tc.expectedCode, tc.expectedMessage)
			base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", tc.expectedStatus, errorResponse.Message))
		})
	}
}

// Negative: удаление из избранного существующего объявления, которого в избранном нет
func (s *TestSuite) TestRemoveAdvertisementNotInFavorites() {
	ad := s.createOwnerAdvertisement()

	base.Step("Покупатель удаляет из избранного объявление, которое туда не добавлял")
	errorResponse := favorites.RemoveFavoriteExpectingError(s.T(), s.buyer.Token, ad.ID, http.StatusNotFound)

	assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeNotFound, apierrors.MessageFavoriteNotFound)
	base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusNotFound, errorResponse.Message))
}

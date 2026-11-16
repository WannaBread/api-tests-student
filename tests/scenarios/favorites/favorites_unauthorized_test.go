package favorites

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

// Negative: все ручки избранного без токена или с невалидным токеном возвращают 401
func (s *TestSuite) TestFavoritesUnauthorized() {
	advertisementID := utils.RandomUUID()
	operations := []struct {
		name string
		call func(t *testing.T, token string) commonModels.ErrorResponse
	}{
		{
			name: "GET /favorites",
			call: func(t *testing.T, token string) commonModels.ErrorResponse {
				return favorites.GetFavoritesExpectingError(t, token, http.StatusUnauthorized)
			},
		},
		{
			name: "POST /favorites",
			call: func(t *testing.T, token string) commonModels.ErrorResponse {
				return favorites.AddFavoriteExpectingError(t, token, advertisementID, http.StatusUnauthorized)
			},
		},
		{
			name: "DELETE /favorites/{ad_id}",
			call: func(t *testing.T, token string) commonModels.ErrorResponse {
				return favorites.RemoveFavoriteExpectingError(t, token, advertisementID, http.StatusUnauthorized)
			},
		},
	}

	for _, operation := range operations {
		for _, tc := range base.UnauthorizedCases() {
			s.Run(operation.name+" "+tc.Name, func() {
				base.Step(fmt.Sprintf("Вызываем %s: %s", operation.name, tc.Name))
				errorResponse := operation.call(s.T(), tc.Token)

				assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeUnauthorized, tc.ExpectedMessage)
				base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusUnauthorized, errorResponse.Message))
			})
		}
	}
}

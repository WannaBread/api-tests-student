package advertisementCreate

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/advertisement"
	base "api-tests-template/tests"
)

// Negative: создание объявления без токена или с невалидным токеном
func (s *TestSuite) TestCreateAdvertisementUnauthorized() {
	for _, tc := range base.UnauthorizedCases() {
		s.Run(tc.Name, func() {
			request := advertisement.NewAdvertisementRequest(base.PhotoPaths(base.DefaultPhotosCount))

			base.Step(fmt.Sprintf("Создаём объявление %q: %s", request.Title, tc.Name))
			errorResponse := advertisement.CreateAdvertisementExpectingError(s.T(), tc.Token, request, http.StatusUnauthorized)

			assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeUnauthorized, tc.ExpectedMessage)
			base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusUnauthorized, errorResponse.Message))
		})
	}
}

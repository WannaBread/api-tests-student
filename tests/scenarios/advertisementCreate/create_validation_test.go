package advertisementCreate

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

// Negative: создание объявления без обязательных параметров или с некорректными значениями
func (s *TestSuite) TestCreateAdvertisementWithInvalidParameters() {
	cases := []struct {
		name            string
		modifyRequest   func(request *models.CreateAdvertisementRequest)
		expectedMessage string
	}{
		{
			name:            "Без title",
			modifyRequest:   func(r *models.CreateAdvertisementRequest) { r.Title = "" },
			expectedMessage: apierrors.MessageTitleRequired,
		},
		{
			name:          "title длиннее максимума",
			modifyRequest: func(r *models.CreateAdvertisementRequest) { r.Title = utils.RandomString(limits.MaxTitleLength + 1) },
			expectedMessage: fmt.Sprintf(apierrors.MessageTitleTooLongTemplate,
				limits.MaxTitleLength, limits.MaxTitleLength+1),
		},
		{
			name:            "Без description",
			modifyRequest:   func(r *models.CreateAdvertisementRequest) { r.Description = "" },
			expectedMessage: apierrors.MessageDescriptionRequired,
		},
		{
			name: "description длиннее максимума",
			modifyRequest: func(r *models.CreateAdvertisementRequest) {
				r.Description = utils.RandomString(limits.MaxDescriptionLength + 1)
			},
			expectedMessage: apierrors.MessageDescriptionLengthInvalid,
		},
		{
			name: "Отрицательная цена",
			modifyRequest: func(r *models.CreateAdvertisementRequest) {
				price := float64(limits.MinPrice - 1)
				r.Price = &price
			},
			expectedMessage: apierrors.MessagePriceOutOfRange,
		},
		{
			name: "Цена выше максимальной",
			modifyRequest: func(r *models.CreateAdvertisementRequest) {
				price := float64(limits.MaxPrice + 1)
				r.Price = &price
			},
			expectedMessage: apierrors.MessagePriceOutOfRange,
		},
		{
			name: "Количество меньше минимального",
			modifyRequest: func(r *models.CreateAdvertisementRequest) {
				quantity := limits.MinQuantity - 1
				r.Quantity = &quantity
			},
			expectedMessage: apierrors.MessageQuantityOutOfRange,
		},
		{
			name: "Количество больше максимального",
			modifyRequest: func(r *models.CreateAdvertisementRequest) {
				quantity := limits.MaxQuantity + 1
				r.Quantity = &quantity
			},
			expectedMessage: apierrors.MessageQuantityOutOfRange,
		},
		{
			name:            "Без фотографий",
			modifyRequest:   func(r *models.CreateAdvertisementRequest) { r.Photos = nil },
			expectedMessage: apierrors.MessagePhotoRequired,
		},
		{
			name: "Фотографий больше максимума",
			modifyRequest: func(r *models.CreateAdvertisementRequest) {
				r.Photos = fixtures.PhotoPaths(limits.MaxPhotosCount + 1)
			},
			expectedMessage: apierrors.MessageTooManyPhotos,
		},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Precondition("Собираем валидный запрос и портим в нём один параметр: " + tc.name)
			request := advertisement.NewAdvertisementRequest(advertisement.DefaultPhotosCount)
			tc.modifyRequest(&request)

			base.Step("Отправляем запрос на создание объявления")
			errorResponse := advertisement.CreateAdvertisementExpectingError(s.T(), s.owner.Token, request, http.StatusBadRequest)

			assertions.ErrorResponseEquals(s.T(), errorResponse, apierrors.CodeInvalidRequest, tc.expectedMessage)
			base.Check(fmt.Sprintf("Получили HTTP %d с ошибкой %q", http.StatusBadRequest, errorResponse.Message))
		})
	}
}

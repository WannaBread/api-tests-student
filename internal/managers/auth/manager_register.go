package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/client/http/auth"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
)

const (
	registeredUserEmailTemplate = "autotest_%s@example.com"
	registeredUserFirstName     = "Autotest"
	randomEmailPartLength       = 12
	randomPasswordLength        = 16
	// passwordComplexitySuffix гарантирует наличие цифры и букв в обоих регистрах в случайном пароле
	passwordComplexitySuffix = "1aA"
)

// RegisterRandomUser регистрирует нового пользователя со случайными email и паролем и возвращает его токен и данные.
// Нужен сценариям, где требуется второй пользователь, например чтобы добавить в избранное чужое объявление.
func RegisterRandomUser(t *testing.T) models.LoginOkResponse {
	t.Helper()
	request := models.RegisterRequest{
		Email:     fmt.Sprintf(registeredUserEmailTemplate, strings.ToLower(utils.RandomString(randomEmailPartLength))),
		Password:  utils.RandomString(randomPasswordLength) + passwordComplexitySuffix,
		FirstName: registeredUserFirstName,
	}
	rawResponse := auth.HttpPostAuthRegister(t, request)
	return response.DecodeJSON[models.LoginOkResponse](t, rawResponse, http.StatusCreated, "Регистрация пользователя")
}

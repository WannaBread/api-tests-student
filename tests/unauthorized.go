package tests

import "api-tests-template/internal/constants/apierrors"

// InvalidToken — заведомо невалидный токен авторизации
const InvalidToken = "incorrect_token"

// UnauthorizedCase описывает запрос без корректной авторизации и ожидаемый текст ошибки 401
type UnauthorizedCase struct {
	Name            string
	Token           string
	ExpectedMessage string
}

// UnauthorizedCases возвращает варианты запроса без авторизации: без токена и с невалидным токеном
func UnauthorizedCases() []UnauthorizedCase {
	return []UnauthorizedCase{
		{Name: "Без токена", Token: "", ExpectedMessage: apierrors.MessageAuthHeaderRequired},
		{Name: "С невалидным токеном", Token: InvalidToken, ExpectedMessage: apierrors.MessageInvalidToken},
	}
}

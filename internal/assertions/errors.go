package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/common/models"
)

// ErrorResponseEquals проверяет код (поле error) и текст (поле message) ошибки в ответе API
func ErrorResponseEquals(t *testing.T, actual models.ErrorResponse, expectedCode, expectedMessage string) {
	t.Helper()
	require.Equal(t, expectedCode, actual.Error, "Неверный код ошибки в поле error")
	require.Equal(t, expectedMessage, actual.Message, "Неверный текст ошибки в поле message")
}

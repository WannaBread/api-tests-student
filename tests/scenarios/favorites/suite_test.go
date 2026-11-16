package favorites

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/managers/advertisement"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

// TestSuite — сценарии работы с избранным (/favorites).
// Своё объявление добавить в избранное нельзя, поэтому участвуют два пользователя:
// owner создаёт объявления, buyer добавляет их в избранное.
type TestSuite struct {
	suite.Suite
	// owner — пользователь из переменных окружения TEST_LOGIN / TEST_PASSWORD
	owner models.LoginOkResponse
	// buyer — новый пользователь, регистрируется на каждый прогон, поэтому его избранное изначально пустое
	buyer models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация владельца объявлений с кредами из переменных окружения")
	s.owner = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	base.Precondition("Регистрация нового пользователя-покупателя для работы с избранным")
	s.buyer = auth.RegisterRandomUser(s.T())
	base.Precondition("Покупатель зарегистрирован: " + s.buyer.User.Email)
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// createOwnerAdvertisement создаёт объявление от имени владельца и удаляет его после завершения теста
func (s *TestSuite) createOwnerAdvertisement() advertisementModels.AdvertisementResponse {
	base.Precondition("Создаём объявление от имени владельца")
	created := advertisement.CreateTemporaryAdvertisement(s.T(), s.owner.Token,
		advertisement.NewAdvertisementRequest(advertisement.DefaultPhotosCount))
	base.Precondition("Объявление создано, id=" + created.ID)
	return created
}

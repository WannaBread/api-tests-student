package advertisementSearch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisement"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

// Название объявления для поиска — случайная строка из трёх частей по searchPartLength символов,
// поэтому каждая часть встречается только в этом объявлении и не зависит от других данных на стенде.
const (
	searchPartLength  = 6
	searchTitleLength = 3 * searchPartLength
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
	searchAd  advertisementModels.AdvertisementResponse
	// Части названия searchAd: начало, середина и конец
	searchBeginning string
	searchMiddle    string
	searchEnd       string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	base.Precondition("Создаём объявление с уникальным названием для поиска")
	request := advertisement.NewAdvertisementRequest(advertisement.DefaultPhotosCount)
	request.Title = utils.RandomString(searchTitleLength)
	s.searchAd = advertisement.CreateTemporaryAdvertisement(s.T(), s.loginData.Token, request)

	s.searchBeginning = request.Title[:searchPartLength]
	s.searchMiddle = request.Title[searchPartLength : 2*searchPartLength]
	s.searchEnd = request.Title[2*searchPartLength:]
	base.Precondition("Объявление создано: " + request.Title)
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	base.Step("Ищем объявления по середине названия: " + s.searchMiddle)
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, s.searchMiddle)
	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})
	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByID(s.T(), list.Items, s.searchAd.ID)
		base.Check("Объявление найдено")
	})
}

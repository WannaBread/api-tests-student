package advertisementSearch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

// Объявление должно заранее существовать на учебном стенде.
const (
	existingAdTitle = "TestTitle"
	searchBeginning = "Test"
	searchMiddle    = "stTi"
	searchEnd       = "Title"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("Проверяем, что константы поиска заданы")
	require.NotEmpty(s.T(), existingAdTitle, "задайте existingAdTitle")
	require.NotEmpty(s.T(), searchBeginning, "задайте searchBeginning")
	require.NotEmpty(s.T(), searchMiddle, "задайте searchMiddle")
	require.NotEmpty(s.T(), searchEnd, "задайте searchEnd")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, searchMiddle)
	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})
	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
	})
}

package advertisementSearch

import (
	"strings"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", s.searchBeginning},
		{"середина", s.searchMiddle},
		{"конец", s.searchEnd},
		{"нижний регистр", strings.ToLower(s.searchMiddle)},
		{"верхний регистр", strings.ToUpper(s.searchMiddle)},
		{"смешанный регистр", utils.MixCase(s.searchMiddle)},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Step("Ищем объявления по подстроке: " + tc.search)
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			assertions.ContainsAdvertisementByID(s.T(), list.Items, s.searchAd.ID)
			base.Check("Объявление найдено")
		})
	}
}

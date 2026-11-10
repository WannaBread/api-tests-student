package utils

import "unicode"

// MixCase чередует верхний и нижний регистр букв, в том числе кириллических.
func MixCase(text string) string {
	runes := []rune(text)
	upper := true
	for i, r := range runes {
		if !unicode.IsLetter(r) {
			continue
		}
		if upper {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
		upper = !upper
	}
	return string(runes)
}

package ceasar

import "unicode/utf8"

func Encode(text string, rotationFactor int) string {
	res := make([]rune, 0, utf8.RuneCountInString(text))
	for _, r := range text {
		if !isLetter(r) {
			res = append(res, r)
		} else {
			res = append(res, getEncodedLetter(r, rotationFactor))
		}
	}

	return string(res)
}

func getEncodedLetter(r rune, factor int) rune {
	if !isLetter(r) {
		return r
	}
	var runeVal int = int(r)
	var lowerVal, upperVal int
	if 'A' <= r && r <= 'Z' {
		lowerVal, upperVal = 'A', 'Z'
	} else {
		lowerVal, upperVal = 'a', 'z'
	}

	for factor > 0 {
		if factor > upperVal-runeVal {
			factor -= upperVal - runeVal + 1
			runeVal = lowerVal
		} else {
			runeVal += factor
			factor = 0
		}
	}

	return rune(runeVal)
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'z'
}

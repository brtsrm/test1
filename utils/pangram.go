package utils

import (
	"strings"
	"unicode"
)

func IsPangramTr(s string) bool {
	return IsPangram(strings.ToLowerSpecial(unicode.TurkishCase, s), "abcçdefgğhıijklmnoöprsştuüvyz")
}
func IsPangramEn(s string) bool {
	return IsPangram(strings.ToLower(s), "abcdefghijklmnopqrstuwxvyz")
}
func IsPangram(s, alphabet string) bool {
	for _, c := range alphabet {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}

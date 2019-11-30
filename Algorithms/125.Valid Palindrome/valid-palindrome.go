package problem125

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	sbs := make([]rune, 0)
	for _, ch := range(s) {
		if (ch >= 'a' && ch <= 'z' || ch >= '1' && ch <= '9') {
			sbs = append(sbs, ch)
		}
	}
	for i := range(sbs) {
		if sbs[i] != sbs[len(sbs)-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	var sbs []rune
	for _, ch := range(s) {
		if (unicode.IsLetter(ch) || unicode.IsDigit(ch)) {		//使用库函数
			sbs = append(sbs, ch)
		}
	}
	for i := range(sbs) {
		if sbs[i] != sbs[len(sbs)-i-1] {
			return false
		}
	}
	return true
}

func isPlindrome3(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !(s[l] <= 'z' && s[l] >= 'a' || s[l] <= '9' && s[l] >= '0') {
			l++
			continue
		}
		if !(s[r] <= 'z' && s[r] >= 'a' || s[r] <= '9' && s[r] >= '0') {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
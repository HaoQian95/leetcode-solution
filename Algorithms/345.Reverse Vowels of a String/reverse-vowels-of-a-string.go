package problem345

func reverseVowels(s string) string {
	sRune := []rune(s)
	l, r := 0, len(sRune)-1
	for l < r {
		if !isVowels(sRune[l]) {
			l++
			continue
		}
		if !isVowels(sRune[r]) {
			r--
			continue
		}
		sRune[l], sRune[r] = sRune[r], sRune[l]
		l++
		r--
	}
	return string(sRune)
}

func isVowels(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

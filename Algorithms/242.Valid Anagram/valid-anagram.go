package problem242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[rune]int)
	mapT := make(map[rune]int) 		//使用int[26]数组效率更好
	for _, ch := range(s) {
		mapS[ch]++
	}
	for _, ch := range(t) {
		mapT[ch]++
	}
	if len(mapS) != len(mapT) {
		return false
	}
	for k, sv := range(mapS) {
		if tv, ok := mapT[k]; !ok || sv != tv {
			return false
		}
	}
	return true
}
package problem205

func isIsomorphic(s string, t string) bool {
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		ch1, ok := ms[s[i]]
		if ok {
			if ch1 != t[i] {
				return false
			}

		} else {
			ms[s[i]] = t[i]
		}
		ch2, ok := mt[t[i]] 
		if ok {						//防止多对一
			if ch2 != s[i] {
				return false
			}
		} else {
			mt[t[i]] = s[i]
		}
	}
	return true
}
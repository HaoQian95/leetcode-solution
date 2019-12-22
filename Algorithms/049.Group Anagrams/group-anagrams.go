package problem049

type runeSlice []rune

func (r runeSlice) Len() int {
	return len(r)
}

func (r runeSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	res := make([][]string, 0)
	
	for _, str := range(strs){
		rs := []rune(str)
        sort.Sort(runeSlice(rs))
		if val, ok := mp[string(rs)]; ok {
			mp[string(rs)] = append(val, str)
		} else {
			mp[string(rs)] = []string{str}
		}
	}

	for _, val := range(mp) {
		res = append(res, val)
	}
	return res
}
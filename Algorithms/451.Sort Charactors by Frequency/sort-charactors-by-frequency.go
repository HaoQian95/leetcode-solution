package problem451

type Pair struct {
	Key rune
	Value int
}
type PairList []Pair

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func frequencySort(s string) string {
	Map := make(map[rune]int)
	res := make([]rune, 0)
	for _, ch := range(s) {
		Map[ch] += 1
	}
	p := make(PairList, len(Map))
	i := 0
	for k, v := range(Map) {
		p[i] = Pair{k, v}
	}
	sort.Sort(p)
	for _, item := range(p) {
		for j := 0; j < item.Value; j++ {
			res = append(res, item.Key)
		}
	}
	return string(res)
}
package problem454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	mp := make(map[int]int)
	
	for i := 0; i < len(A); i++ {
		for j := 0; j <len(B); j++ {
			mp[A[i]+B[j]] += 1
		}
	}

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			target := 0 - C[i] - D[j]
			if val, ok := mp[target]; ok {
				res += val 
			}
		}
	}
	return res
}
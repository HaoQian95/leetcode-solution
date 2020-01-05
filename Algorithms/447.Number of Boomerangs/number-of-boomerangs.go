package problem447

//用 distance 做 map 的 key 值会不准确（float64计算有误差），所以不开方 
func numberOfBoomerangs(points [][]int) int {
	res := 0
	length := len(points)
	mpS := make([]map[int]int, length)

	for i := 0; i < length; i++ {
		for j := i+1; j < length; j++ {
			distance2pow := (points[i][0]-points[j][0]) * (points[i][0]-points[j][0]) + (points[i][1]-points[j][1]) * (points[i][1]-points[j][1])
			if mpS[i] == nil {
				mpS[i] = make(map[int]int)
			}
			if mpS[j] == nil {
				mpS[j] = make(map[int]int)
			}
			mpS[i][distance2pow] += 1
			mpS[j][distance2pow] += 1
		}
	}

	for i := 0; i < length; i++ {
		for _, v := range(mpS[i]) {
			res += v * (v - 1)
		}
	}
	return res
}
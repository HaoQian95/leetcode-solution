package problem149

func maxPoints(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		m := make(map[float64]int)
		vtl := 0	//和points[i]在一条竖线上的点的数量	
		dup := 1	//和points[i]重合的点的数量
		for j := 0; j <len(points); j++ {
			if i == j {
				continue
			}
			if points[i][0] == points[j][0] {
				if points[i][1] == points[j][1] {
					dup++
				} else {
					vtl++
				}
			} else {
				k := float64(points[i][1]-points[j][1]) / float64(points[i][0]-points[j][0])
				m[float64(k)]++
			}
		}
		max := vtl 
		for _, v := range(m) {
			if max < v {
				max = v
			}
		}
		if res < (max + dup) {
			res = max + dup
		}
	}

	return res
}
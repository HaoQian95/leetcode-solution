package problem149

//浮点数精度问题 [[0,0],[94911151,94911150],[94911152,94911151]] 测试用例通不过
func maxPoints(points [][]int) int {
	res := 0
	if len(points) <= 2 {
		return len(points)
	}
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
				k := float64((points[i][1]-points[j][1])) / float64((points[i][0]-points[j][0]))
				m[k]++
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

//解决办法：不算出斜率的最后结果，将约分后的最简形式转化为string，作为 map 的 key;
func maxPoints2(points [][]int) int {
	res := 0
	if len(points) <= 2 {
		return len(points)
	}

	for i := 0; i < len(points); i++ {
		m := make(map[string]int)
		vtl := 0 //和points[i]在一条竖线上的点的数量
		dup := 1 //和points[i]重合的点的数量
		for j := i+1; j <len(points); j++ { //j 可以从 i+1 开始，因为 j 和 i 之前（包括i点）连成的直线已经计算过了
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
				dy := points[i][1] - points[j][1]
				dx := points[i][0] - points[j][0]
				g := gcd(dx, dy)
				s := strconv.Itoa(dy/g) + "_" + strconv.Itoa(dx/g)
				m[s]++
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

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}
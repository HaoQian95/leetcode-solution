package problem070

var Map map[int]int = make(map[int]int)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, ok := Map[n]; ok != false {
		return val
	}
	Map[n] = climbStairs(n-1) + climbStairs(n-2)
	return Map[n]
}



//Time limit exceeded
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}


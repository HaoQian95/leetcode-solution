package problem011

//O(n)
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		var  currentArea int
		if height[l] < height[r] {
			currentArea = (r - l) * height[l]
			l++
		} else {
			currentArea = (r - l) * height[r]
			r--
		}
		if currentArea > res {
			res = currentArea
		}
	}
	return res
} 

//暴力解法 O(n^2)
func maxArea2(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i+1; j < len(height); j++ {
			var temp int
			if height[i] < height[j] {
				temp = (j - i) * height[i]
			} else {
				temp = (j - i) * height[j]
			}
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
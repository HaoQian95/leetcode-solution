package problem209

//滑动窗口 O(n)
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
        return 0
    }
    l, r := 0, 0
	currentSum := nums[l]
	res := len(nums)+1
	for r < len(nums) && l <= r{
		if currentSum >= s {
			if res > r-l+1 {
				res = r-l+1
			}
			currentSum -= nums[l]
			l++
		} else {
            if r == len(nums)-1 {
                break
            }
			currentSum += nums[r+1]
			r++
		}
	}
    if res == len(nums)+1{
        return 0
    }
	return res
}
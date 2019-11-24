package problem075

//两次遍历
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	var numsZero, numsOne, numsTwo int
	for _, val := range(nums) {
		if val == 0{
			numsZero++
		} else if val == 1 {
			numsOne++
		} else {
			numsTwo++
		}
	}
	for i := range(nums) {
		if i < numsZero {
			nums[i] = 0
		} else if i >= numsZero + numsOne {
			nums[i] = 2
		} else {
			nums[i] = 1
		}
	}
}

//一次遍历（三路快排）
func sortColors2(nums []int) {
	it, i, gt := -1, 0, len(nums)
	for i < gt {
		if nums[i] == 0 {
			it++
			nums[i], nums[it] = nums[it], nums[i]
			i++
		} else if nums[i] == 2 {
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		} else {
			i++
		}
	}
}
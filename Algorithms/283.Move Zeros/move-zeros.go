package problem283

func moveZeroes(nums []int)  {
	position := 0
	for _, num := range(nums) {
		if num != 0 {
			nums[position] = num
			position++
		}
	}
	for position < len(nums) {
		nums[position] = 0
		position++
	}
}
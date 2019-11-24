package problem026

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	position := 0
	for i := position+1; i < len(nums); i++ {
		if nums[i] != nums[position] {
			position++
			nums[position] = nums[i]
		}
	}
	return position+1
}
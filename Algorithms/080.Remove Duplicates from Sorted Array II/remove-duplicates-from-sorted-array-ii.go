package problem080

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	position := 0
	count := 1	//记录重复元素的个数
	for i := position+1; i < len(nums); i++ {
		if nums[i] != nums[position] {
			position++
			nums[position] = nums[i]
			count = 1
		} else if count == 1{
			position++
			nums[position] = nums[i]
			count++
		}
	}
	return position+1
}
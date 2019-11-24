package problem027

func removeElement(nums []int, val int) int {
	position := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[position] = nums[i]
			position++
		}
	}
	return position
}
package problem217

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range(nums) {
		if set[v] {
			return false
		}
		set[v] = true
	}
	return true
}
package problem001

func twoSum(nums []int, target int) []int {
	Map := make(map[int]int)
	for k, v := range(nums) {
		t := target - v
		if item, ok := Map[t]; ok {
			return []int{k, item}
		} else {
			Map[v] = k
		}
	}
	return []int{0, 0}
}
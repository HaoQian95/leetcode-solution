package problem349

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)
	for _, v1 := range(nums1) {
		set[v1] = true
	}
	for _, v2 := range(nums2) {
		if set[v2] {
			res = append(res, v2)
			delete(set, v2)
			//set[v2] = false
		}
	}
	return res
}
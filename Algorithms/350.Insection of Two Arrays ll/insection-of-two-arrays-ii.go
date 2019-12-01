package problem350

func intersect(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	res := make([]int, 0)
	for _, v := range(nums2) {
		dict[v] += 1
	}
	for _, v2 := range(nums1) {
		if dict[v2] > 0 {
			res = append(res, v2)
			dict[v2] -= 1
		}
	}
	return res
}
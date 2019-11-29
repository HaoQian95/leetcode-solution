package problem088

//尾插法
func merge(nums1 []int, m int, nums2 []int, n int) {
	count := m+n-1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[count] = nums1[m]
			m--
		} else {
			nums1[count] = nums2[n]
			n--
		}
		count--
	}
	for n >= 0 {
		nums1[count] = nums2[n]
		count--
		n--
	}
}
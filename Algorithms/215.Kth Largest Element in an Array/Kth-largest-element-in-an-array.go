package problem215

//快排思想，最坏情况时间复杂度On^2
func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums)-1
	for left < right {
		location := partition(nums, left, right)
		if right-location+1 == k {
			return nums[location]
		}else if right-location+1 < k {
			k = k - (right-location+1)
			right = location-1
		}else {
			left = location+1
		}
	}
	return nums[left]
}


func partition(nums []int, left int, right int) int {
	temp := nums[left]
	for left < right {
		for nums[right] >= temp && left < right{
			right--
		}
		nums[left] = nums[right]
		for nums[left] <= temp && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = temp
	return left
}

//暴力解法, 时间复杂度
func findKthLargest2(nums []int, k int) int {
	for i := 0; i < k; i++ {
		temp := nums[0]
		location := 0
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > temp {
				location = j
				temp = nums[j]
			}
		}
		nums[len(nums)-i-1], nums[location] = nums[location], nums[len(nums)-i-1]
	}
	return nums[len(nums)-k]
}
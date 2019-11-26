package problem215

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
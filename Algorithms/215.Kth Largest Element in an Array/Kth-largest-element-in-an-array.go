package problem215

func findKthLargest(nums []int, k int) int {

}


func partition(nums []int, left int, right int) int {
	temp := nums[left]
	for left < right {
		for nums[right] > temp && left < right{
			right--
		}
		nums[left] = nums[right]
		for nums[left] < temp && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = temp
	return left
}
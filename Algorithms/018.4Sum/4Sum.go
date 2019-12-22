package problem018

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
		for j := i+1; j < len(nums)-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
			subTarget := target - nums[i] - nums[j]
			for left, right := j+1, len(nums)-1; left < right; {
				if nums[left] + nums[right] < subTarget {
					left++
				} else if nums[left] + nums[right] > subTarget {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}
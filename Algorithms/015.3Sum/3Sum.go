package problem015

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
    sort.Ints(nums)
    
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }
		target := 0 - nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[j] + nums[k] < target {
				j++
			} else if nums[k] + nums[j] > target {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return res
}
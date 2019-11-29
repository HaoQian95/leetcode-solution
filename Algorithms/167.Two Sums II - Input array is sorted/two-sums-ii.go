package problem167

//暴力解法
func twoSum(numbers []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		tar := target - numbers[i]
		for j := i+1; j < len(numbers); j++ {
			if numbers[j] > tar {
				break
			}else if tar == numbers[j] {
				return append(res, i+1, j+1)
			}
		}
	}
	return res
}

//对撞指针 O(n)
func twoSum2(numbers []int, target int) []int {
	res := make([]int, 0)
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i] + numbers[j] < target{
			i++
		} else if numbers[i] +numbers[j] == target {
			return append(res, i+1, j+1)
		} else {
			j--
		}
	}
	return res
}

//二分查找 O(nlogn)
func twoSum3(numbers []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		tar := target - numbers[i]
		index, ok := binarySearch(numbers, i+1, tar)
		if ok {
			return append(res, i+1, index+1)
		}
	}
	return res
}

func binarySearch(numbers []int, start int, target int) (int, bool) {
	end := len(numbers)-1
	for start <= end {
		mid := start + (end-start)/2
		if numbers[mid] < target {
			start = mid + 1
		} else if numbers[mid] > target {
			end = mid - 1
		} else {
			return mid, true
		}
	}
	return -1, false
}
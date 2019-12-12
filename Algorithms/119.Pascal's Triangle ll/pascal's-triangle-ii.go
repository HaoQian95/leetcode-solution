package problem119

//递归
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := make([]int, rowIndex+1)
	lastRes := getRow(rowIndex-1)
	res[0], res[rowIndex] = 1, 1
	for i := 1; i < rowIndex; i++ {
		res[i] = lastRes[i-1] + lastRes[i]
	}
	return res
}
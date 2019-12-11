package problem118

//递归
func generate(numRows int) [][]int {
	if numRows == 0 {
		return make([][]int, 0)
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	lastRow := make([]int, numRows)
	lastRes := generate(numRows-1)
	lastRow[0] = 1
	lastRow[numRows-1] = 1
	for i := 1; i < numRows-1; i++ {
		lastRow[i] = lastRes[numRows-2][i-1] + lastRes[numRows-2][i]
	}
	return append(lastRes, lastRow)
}
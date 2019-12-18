package problem096

//memorization(否则会 time limit)
func numTrees (n int) int {
	return numSubTrees(1, n)
}

var Map map[int]int = make(map[int]int)
 
func numSubTrees(left int, right int) int {
	if left == right {
		return 1
	}
	if left > right {
		return 0
	}
	if val, ok := Map[right-left+1]; ok != false {
		return val
	}
	res := 0
	for i := left; i <= right; i++ {
		leftSubTrees := numSubTrees(left, i-1)
		rightSubTrees := numSubTrees(i+1, right)
		if leftSubTrees == 0 {
			res += rightSubTrees
		} else if rightSubTrees == 0 {
			res += leftSubTrees
		} else {
			res += leftSubTrees * rightSubTrees
		}
	}
	Map[right-left+1] = res
	return res
}
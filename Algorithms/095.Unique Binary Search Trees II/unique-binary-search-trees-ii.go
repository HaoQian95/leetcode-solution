package problem095

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n < 1 {
		return nil
	}
	return generateSubTrees(1, n)
}

func generateSubTrees(left int, right int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if right < left {
		res = append(res, nil)
		return res
	}
	if left == right {
		node := TreeNode{Val: left}
		res = append(res, &node)
		return res
	}

	for i := left; i <= right; i++ {
		l := generateSubTrees(left, i-1)
		r := generateSubTrees(i+1, right)
		for j := 0; j < len(l); j++ {
			for k := 0; k < len(r); k++ {
				node := TreeNode{Val : i}
				node.Left = l[j]
				node.Right = r[k]
				res = append(res, &node)
			}
		}
	}
	return res
}
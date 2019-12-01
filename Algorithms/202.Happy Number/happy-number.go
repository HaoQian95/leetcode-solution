package problem202

//递归：测试不通过。应该是测试代码跟全局变量定义在一个包中，导致全局变量不更新
var set map[int]bool = make(map[int]bool)

func isHappy(n int) bool {
    if set[n] {
        return false
    }
    set[n] = true
    res := 0
	for n != 0 {
		res += (n%10)*(n%10)
		n = n/10
	}
	if res == 1 {
		return true
	}
    //fmt.Println(res)
	return isHappy(res)
}

//改用迭代
func isHappy2(n int) bool {
	set := make(map[int]bool)
	for {
		if set[n] {
			return false
		}
		set[n] = true
		res := 0
		for n != 0 {
			res += (n%10)*(n%10)
			n = n/10
		}
		if res == 1 {
			break
		}
		n = res
	 }
	 return true
}
package problem050

//O(logn)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	var abs_n int
	var res float64
	if n < 0 {
		abs_n = 0 - n
	} else {
		abs_n = n
	}
	temp := myPow(x, abs_n / 2)
	if abs_n % 2 == 0 {
		res = temp * temp
	} else {
		res = x * temp * temp
	}
	if n < 0 {
		res = 1 / res
	}
	return res
	/*改良
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	b := myPow(x, n / 2)
	p := myPow(x, n % 2)
	return b * b * p
	*/
}

//Time limit exceeded : O(n)
func myPow2(x float64, n int) float64 {
	var flag bool
	var res float64
	res = 1
	if n < 0 {
		flag = true
		n = 0 - n
	}
	for i := 0; i < n; i++ {
		res *= x
	}
	if flag {
		return 1/res
	}
	return res
}
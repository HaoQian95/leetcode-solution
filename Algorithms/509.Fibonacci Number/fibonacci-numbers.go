package problem509

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

//Memorization
var Map map[int]int = make(map[int]int)
func fib2(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	if val, ok := Map[N]; ok != false {
		return val
	}
	Map[N] = fib2(N-1) + fib2(N-2)
	return Map[N]
}
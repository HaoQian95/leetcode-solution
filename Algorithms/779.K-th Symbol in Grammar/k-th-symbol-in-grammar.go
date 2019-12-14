package problem779

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	lastInt := kthGrammar(N-1, (K+1)/2)
	if K % 2 == 1 {			//0 -> 01	1 -> 10
		return lastInt
	}
	return (lastInt + 1) % 2
} 
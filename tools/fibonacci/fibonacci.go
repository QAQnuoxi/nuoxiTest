package fibonacci

func Fibonacci(n int) int {
	a, b := n%2, 1
	for i := 0; i < n/2; i++ {
		a += b
		b += a
	}
	return a
}

package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	//预先定义的一组斐波那契数列作为测试用例
	fcMap := map[int]int{}
	fcMap[0] = 0
	fcMap[1] = 1
	fcMap[2] = 1
	fcMap[3] = 2
	fcMap[4] = 3
	fcMap[5] = 5
	fcMap[6] = 8
	fcMap[7] = 13
	fcMap[8] = 21
	fcMap[9] = 34
	for k, v := range fcMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)
		}
	}
}
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(66)
	}
}

package q3

import "testing"
//性能测试函数

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i< b.N; i++{
		GetPrimes(1000)
	}
}
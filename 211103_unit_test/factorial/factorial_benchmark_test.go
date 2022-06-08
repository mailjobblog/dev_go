package factorial

import (
	"testing"
)

// 基准测试
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

// 基准测试之性能比较测试
func benchmarkFact(b *testing.B, x int) {
	for i := 0; i < b.N; i++ {
		Factorial(x)
	}
}
func BenchmarkFact1(b *testing.B)   { benchmarkFact(b, 1) }
func BenchmarkFact5(b *testing.B)   { benchmarkFact(b, 5) }
func BenchmarkFact10(b *testing.B)  { benchmarkFact(b, 10) }
func BenchmarkFact50(b *testing.B)  { benchmarkFact(b, 50) }
func BenchmarkFact100(b *testing.B) { benchmarkFact(b, 100) }

// 基准测试之并行测试
func BenchmarkFactorialParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Factorial(10)
		}
	})
}

package race

import "testing"

func BenchmarkAtomicWith10G(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicIncrement(10)
	}
}

func BenchmarkAtomicWith2G(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicIncrement(2)
	}
}

func BenchmarkMutexWith10G(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutexIncrement(10)
	}
}

func BenchmarkMutexWith2G(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutexIncrement(2)
	}
}

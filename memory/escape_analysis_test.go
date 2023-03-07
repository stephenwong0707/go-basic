package memory

import "testing"

func BenchmarkPtrInPtrOut(b *testing.B) {
	var tmp *bar
	for i := 0; i < b.N; i++ {
		tmp = ptrInPtrOut(&input)
		output = *tmp
	}
}

// 0 allocs/op
func BenchmarkPtrInValOut(b *testing.B) {
	var tmp bar
	for i := 0; i < b.N; i++ {
		tmp = ptrInValOut(&input)
		output = tmp
	}
}

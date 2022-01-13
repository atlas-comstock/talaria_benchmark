package main

import "testing"

func BenchmarkTrim(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		trim()
	}
}
func BenchmarkLuaTrim(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		luaTrim()
	}
}

package main

import "testing"

func TestFoo(t *testing.T) {
	x := foo(2, 3)
	if x != 5 {

		t.Error("expected", 5, " got", x)
	}

}

func BenchmarkFoo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		foo(2, 3, 4)
	}

}

package main

// int rand(int i) { return i+2; }
import "C"

import "testing"

var resGo int
var resC C.int

func BenchmarkCgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resC = C.rand(resC)
	}
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resGo = rand(resGo)
	}
}

func rand(i int) int {
	return i + 2
}

func main() {
	testing.Main(
		func(string, string) (bool, error) {
			return true, nil
		},
		nil,
		[]testing.InternalBenchmark{
			{"BenchmarkCgo", BenchmarkCgo},
			{"BenchmarkGo", BenchmarkGo},
		},
		nil,
	)
	//fmt.Println(res)
}

the Benchamerk checks an overhead of making cgo call compared to Go function call.
for golang 1.9.1 the penalty is about 40x 

```
$ go version
go version go1.9.1 linux/amd64

$ go run main.go -test.bench=.
goos: linux
goarch: amd64
BenchmarkCgo-4   	20000000	        62.9 ns/op
BenchmarkGo-4    	2000000000	         1.54 ns/op

```

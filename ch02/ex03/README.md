ループを使う形に書き直したら4倍以上遅くなった。


❯ go test -bench=PopCount$
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex03
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount-8     1000000000               0.2832 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex03    0.415s

❯ go test -bench=PopCount2
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex03
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount2-8    325767806                3.649 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex03    1.757s
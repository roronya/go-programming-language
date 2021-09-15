テストデータには、２進数表記にするとすべて1が立っているようなuint64の値を用いた。
各PopCountをN回繰り返し、1度のPopCountの実行にかかった平均値で比較する。

最初のバージョンと比較して約100倍、PopCount2と比較して約8倍遅くなった。

❯ go test -bench=PopCount3
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex04
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount3-8    47324504                25.34 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex04    1.325s

下にex03で行ったベンチマークの結果も示す。

❯ go test -bench=PopCount$
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex03
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount-8     1000000000               0.2831 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex03    0.412s

❯ go test -bench=PopCount2
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex03
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount2-8    326928751                3.616 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex03    1.659s


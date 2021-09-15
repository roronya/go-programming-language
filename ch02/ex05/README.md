テストデータには、２進数表記にするとすべて1が立っているようなuint64の値を用いた。
各PopCountをN回繰り返し、1度のPopCountの実行にかかった平均値で比較する。

他の手法に比べ、初期のPopCountの約150倍、PopCount2の約14倍、PopCount3の約1.5倍遅くなっている。

❯ go test -bench=PopCount4
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex05
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount4-8    28252220                42.31 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex05    1.338s

PopCount4はbit列の中で1が立っている数だけループを行うような手法なので、すべて1が立っている場合は最悪実行時間になる。
しかし、通常は入力値のbit列のすべてが1というケースは稀であるので、1が立っている箇所がまばらな平均的な値で性能も見ても良いかもしれない。

平均的な値とは、この場合、64bitのうち半分の32bitが1が立っているようなデータである。
そのようなデータで性能を見ると、他の手法に比べ、初期のPopCountの約50倍、PopCount2の約4倍遅くなり、PopCount3の約2倍早くなっている。
(他の手法は入力値に現れる1の数には依存しない手法なので、単純に比較できる。)

❯ go test -bench=PopCount4_2
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch02/ex05
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkPopCount4_2-8          82337580                12.98 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch02/ex05    1.204s


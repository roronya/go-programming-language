package main

import (
	"strings"
	"testing"
)

func echo02(a []string) {
	s, sep := "", ""
	for _, arg := range a {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func echo03(a []string) {
	strings.Join(a, " ")
	// fmt.Println(strings.Join(a, " "))
}

func BenchmarkEcho02(b *testing.B) {
	a := make([]string, 100000)
	for i := 0; i < b.N; i++ {
		echo02(a)
	}

}
func BenchmarkEcho03(b *testing.B) {
	a := make([]string, 100000)
	for i := 0; i < b.N; i++ {
		echo03(a)
	}

}

/**
引数が10万個渡されたときのベンチマークを取った
echo02では1回の呼び出しが平均して約597msだったのに対してecho03は0.7msだった

❯ go test -bench=Echo02
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch01/ex03
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkEcho02-8              2         597911712 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch01/ex03    2.006s

❯ go test -bench=Echo03
goos: darwin
goarch: amd64
pkg: github.com/roronya/go-programming-language/ch01/ex03
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkEcho03-8           1458            706515 ns/op
PASS
ok      github.com/roronya/go-programming-language/ch01/ex03    1.221s

*/

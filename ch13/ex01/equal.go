package equal

import (
	"math"
	"math/cmplx"
	"reflect"
)

// 10^-9よりも小さい違いしかなければ等しいとみなす比較関数
// 一桁ずつみて10^-9桁で打ち切る
func Equal(x, y interface{}) bool {
	return equal(reflect.ValueOf(x), reflect.ValueOf(y))
}

func equal(x, y reflect.Value) bool {
	if x.Type() != y.Type() {
		return false
	}

	// 数値型以外ならfalseを返す
	// intなら普通に比較する
	// floatなら整数部と小数部に分けてそれぞれ比較
	// 小数部は10^9倍して整数部を比較すれば良い(桁あふれしないか？)
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int() == y.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return x.Uint() == y.Uint()
	case reflect.Float32, reflect.Float64:
		return math.Abs(x.Float()-y.Float()) < 1e-9
	case reflect.Complex64, reflect.Complex128:
		return cmplx.Abs(x.Complex()-y.Complex()) < 1e-9
	default:
		return false
	}
}

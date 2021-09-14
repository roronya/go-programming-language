package popcount

/**
> 8ビットの値が取り得る結果すべてのテーブルpcを事前に計算する
変数pcは添字の整数を2進数表記したときに現れる1の数を返す
*/
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		// pc[0] = pc[0] + byte(00000000&00000001) = 00000000 + 00000000 = 00000000 = 0
		// pc[1] = pc[0] + byte(00000001&00000001) = 00000000 + 00000001 = 00000001 = 1
		// pc[2] = pc[1] + byte(00000010&00000001) = 00000001 + 00000000 = 00000001 = 1
		// pc[3] = pc[1] + byte(00000011&00000001) = 00000001 + 00000001 = 00000010 = 2
		// pc[4] = pc[2] + byte(00000100&00000001) = 00000001 + 00000000 = 00000001 = 1
		// pc[5] = pc[2] + byte(00000101&00000001) = 00000001 + 00000001 = 00000010 = 2
		// ...
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int {
	// xを8bitずつ千切って、それぞれの1が立っている数を数えている
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

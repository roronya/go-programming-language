package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func printReflectValue(name string, x reflect.Value) {
	fmt.Printf("--------------\nname: %s\ntype: %s\nvalue: %v\n", name, x.Type(), x)
}

func main() {
	// 基本的な振る舞い
	x := 2
	printReflectValue("x=2; reflect.ValueOf(x)", reflect.ValueOf(x))                                  // => type:int value:2
	printReflectValue("x=2; reflect.ValueOf(&x)", reflect.ValueOf(&x))                                // => type:*int value: 0xc000..
	printReflectValue("x=2; reflect.ValueOf(&x).Elem()", reflect.ValueOf(&x).Elem())                  // => type:int value:2
	printReflectValue("x=2; reflect.ValueOf(&x).Elem()", reflect.ValueOf(reflect.ValueOf(&x).Elem())) // => type:int value:2

	// interfaceに対する動き
	var s io.Reader
	// io.Readerは型であって値ではないのでreflect.ValueOfは呼べない
	// printReflectValue("io.Reader", reflect.ValueOf(io.Reader))

	// panicになる
	// panic: reflect: call of reflect.Value.Type on zero Value
	// sは何も代入していないので、type:nil value:nilなインタフェース値になっているためだと覆う
	// printReflectValue("s", reflect.ValueOf(s))

	// sに具象値を代入する
	s = os.Stdout
	printReflectValue("var s io.Reader; s=os.Stdout; reflect.ValueOf(s)", reflect.ValueOf(s))                                             // type:*io.File value: 0x000...
	printReflectValue("var s io.Reader; s=os.Stdout; reflect.ValueOf(&s)", reflect.ValueOf(&s))                                           // type:*io.Reader value: 0x000...
	printReflectValue("var s io.Reader; s=os.Stdout; reflect.ValueOf(&s).Elem()", reflect.ValueOf(&s).Elem())                             // type:io.Reader value: 0x000...
	printReflectValue("var s io.Reader; s=os.Stdout; reflect.ValueOf(&s).Elem().Elem()", reflect.ValueOf(&s).Elem().Elem())               // type:*io.File value: 0x000...
	printReflectValue("var s io.Reader; s=os.Stdout; reflect.ValueOf(&s).Elem().Elem().Elem()", reflect.ValueOf(&s).Elem().Elem().Elem()) // type:*io.File value: 0x000...
	/**
	sはインターフェースであるos.Reader型である
	reflect.ValueOfはインタフェースに対して呼ぶと、呼び出されたタイミングで設定されている動的な型と動的な値を取り出す
	sに対するreflect.ValueOfはそのタイミングで設定されているos.Stdoutの型と値を返す(*os.Fileと&{0xc000...})

	一方、sのポインタに対してreflect.ValueOfを呼び出すとインタフェース型で包まれている様子が見れる
	その結果にElem()を呼び出すと更に構造が以下のようになっていることがわかる
	{
		type: *io.Reader,
		value: 0xc000...
		elem: {
			type: io.Reader,
			value: 0xc000...
			elem: {
				type: *os.File,
				value: &{0xc000...}
				elem: {
					type: os.File,
					value: &{0xc000...}
				}
			}
		}
	}
	*/

	/**
	上の例を簡略化してみる
	interface{}な値のポインタに対してreflect.ValueOfを呼んでElemで深ぼっていくと以下のような構造になっている
	Typeの構造:  *interface{} =Elem()=> interface =Elem()=> int
	Valueの構造: 0xc000...    =Elem()=> 3         =Elem()=> 3
	変数に代入した値の型にたどり着くには2階層Elem()を呼ぶ必要があるが、
	Valueは1度Elem()を呼び出せばinterface{}型ではあるが取り出せる
	*/
	var i interface{} = 3
	printReflectValue("var i interface{}=3; reflect.ValueOf(i)", reflect.ValueOf(i))
	printReflectValue("var i interface{}=3; reflect.ValueOf(&i)", reflect.ValueOf(&i))
	printReflectValue("var i interface{}=3; reflect.ValueOf(&i).Elem()", reflect.ValueOf(&i).Elem())
	printReflectValue("var i interface{}=3; reflect.ValueOf(&i).Elem().Elem()", reflect.ValueOf(&i).Elem().Elem())

	/**
	値の書き換え
	*/
	x = 2
	d := reflect.ValueOf(&x).Elem() // dは変数xを参照している
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x) // => 3
	/**
	L81は、reflect.ValueOf(&x).Elem()という処理により *interface=>interface の構造を辿って d にinterface{}型の2を代入している
	d.Addr().Interface().(*int)
	*/

	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d = c.Elem()
	e := c
	fmt.Println(a.CanAddr()) // false
	fmt.Println(b.CanAddr()) // false
	fmt.Println(c.CanAddr()) // false
	fmt.Println(d.CanAddr()) // true
	fmt.Println(e.CanAddr()) // false
	printReflectValue("a", a)
	printReflectValue("b", b)
	printReflectValue("c", c)
	printReflectValue("d", d)
	printReflectValue("e", e)
}

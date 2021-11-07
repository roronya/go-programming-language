package main

import (
	"bytes"

	"golang.org/x/net/html"
)

func ExampleOutline() {
	in, _ := html.Parse(bytes.NewBufferString("<p>this is text node</p>"))

	// 最初からdepthを2に設定しているので、ローカル変数を共有していればこのdepthが使われてインデントされた状態で出力され
	depth = 2
	outline(in)
	// Output:
	//     <html>
	//       <head>
	//       </head>
	//       <body>
	//         <p>
	//         </p>
	//       </body>
	//     </html>
}

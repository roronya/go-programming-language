package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	f, _ := os.Open("test.html")
	txt, _ := ioutil.ReadAll(f)
	doc, _ := html.Parse(bytes.NewBuffer(txt))

	out = new(bytes.Buffer)
	forEachNode(doc, startElement, endElement)

	_, err := html.Parse(out.(*bytes.Buffer))
	if err != nil {
		log.Fatalf("html parse error!\nerror:%s\ninput:%s", err, out.(*bytes.Buffer).String())
	}
}

func ExampleCommentNode() {
	doc, _ := html.Parse(bytes.NewBufferString("<!-- this is comment -->"))
	forEachNode(doc, startElement, endElement)
	// Output:
	// <!-- this is comment -->
	// <html>
	//   <head/>
	//   <body/>
	// </html>
}

func ExampleTextNode() {
	doc, _ := html.Parse(bytes.NewBufferString("<p>this is text node</p>"))
	forEachNode(doc, startElement, endElement)
	// Output:
	// <html>
	//   <head/>
	//   <body>
	//     <p>
	//       this is text node
	//     </p>
	//   </body>
	// </html>
}

func ExampleAttr() {
	doc, _ := html.Parse(bytes.NewBufferString("<img href=\"link\" alt=\"this is alt\"/>"))
	forEachNode(doc, startElement, endElement)
	// Output:
	// <html>
	//   <head/>
	//   <body>
	//     <img href="link" alt="this is alt"/>
	//   </body>
	// </html>
}

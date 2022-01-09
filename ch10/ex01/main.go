package main

import (
	"flag"
	"image"
	"image/jpeg"
	_ "image/png" // PNGデコーダを登録する
	"io"
	"log"
	"os"
)

/**
& cat input.jpg |  go run ./main.go -out png
2022/01/09 09:27:09 unsupported output format: png
$ cat input.png |  go run ./main.go -out png
2022/01/09 09:28:11 unsupported output format: png
*/
func main() {
	var outf string
	flag.StringVar(&outf, "out", "jpeg", "output format")
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if !(kind == "jpeg" || kind == "png") {
		log.Fatalf("unsupported input format: %s\n", kind)
	}

	switch outf {
	case "jpeg":
		err = toJPEG(img, os.Stdout)
	default:
		log.Fatalf("unsupported output format: %s\n", outf)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

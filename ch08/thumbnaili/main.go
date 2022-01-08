package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Image returns a thumbnail-size version of src.
func Image(src image.Image) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// writes a thumbnail-size version of it to w.
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes
// a thumbnail-size version of it to outfile.
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g. "foo.thumb.jpeg".
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG"
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 注意: 正しくない
// ゴルーチンを待たずに呼び出し元に返るのでImageFileの処理が完了しない
// 実際に実行するとthumbnailが作られない
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f) // エラーを無視
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			if _, err := ImageFile(f); err != nil {
				log.Println(err)
			}
			ch <- struct{}{}
		}(f)
	}

	// ゴルーチンの完了を待つ
	for range filenames {
		<-ch
	}
}

// 最初のエラーを見つけると呼び出し元に返す
// errorsを空にするチャネルがいないから、チャネルに値を送ろうとして待たされる
// この例だとmainルーチンがすぐに終了するので、プログラムとしては終了するが、チャネルの処理は詰まっている（？）
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	// ゴルーチンの完了を待つ
	for range filenames {
		if err := <-errors; err != nil {
			return err // 注意: 正しくない: ゴルーチンのリーク
		}
	}
	return nil
}

// ch := make(chan item, len(filenames))
// で十分なバッファのチャネルを作るから、チャネルの処理はつまったりはしない
// チャネルを空にする処理がないことが問題ではなくて、チャネルの処理につまるのが問題っぽい
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64, 4)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)

		go func(f string) {
			fmt.Println(f)
			sizes <- 1
			wg.Done()
		}(f)
		/**
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
		*/
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func reader(ch chan string) {
	filenames := []string{
		"a.jpeg",
		"b.jpeg",
		"c.jpeg",
		"d.jpeg",
	}
	for _, f := range filenames {
		ch <- f
	}
}

func main() {
	ch := make(chan string)
	go reader(ch)
	go makeThumbnails6(ch)
}

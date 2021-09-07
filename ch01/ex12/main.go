package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		var cycles, res, size, nframes, delay = 5, 0.001, 100, 64, 8
		f := r.Form
		var err error
		if v := f.Get("cycles"); v != "" {
			cycles, err = strconv.Atoi(v)
		}
		if v := f.Get("res"); v != "" {
			res, err = strconv.ParseFloat(v, 64)
		}
		if v := f.Get("size"); v != "" {
			size, err = strconv.Atoi(v)
		}
		if v := f.Get("nframes"); v != "" {
			nframes, err = strconv.Atoi(v)
		}
		if v := f.Get("delay"); v != "" {
			delay, err = strconv.Atoi(v)
		}
		if err != nil {
			fmt.Fprintf(w, "query error")
		}
		lissajous(w, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}

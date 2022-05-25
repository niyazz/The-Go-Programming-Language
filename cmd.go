package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"io"
	"image/gif"
	"math"
	"math/rand"
	"time"
)

var green = color.RGBA{0x12, 0x80, 0x12, 0xff}
var red = color.RGBA{0x90, 0x10, 0x12, 0xff}
var blue = color.RGBA{0x12, 0x10, 0x90, 0xff}
var palette = []color.Color{color.Black, green, red, blue}
var mu sync.Mutex
var count int

func main(){
	handler := func(w http.ResponseWriter, r *http.Request){
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer){
	const(
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			
// Task 1.6: go run main.go > out.gif
// Измените программу lissajous так, чтобы она генерировала 
// изображения разных цветов, добавляя в палитру palette больше значений, 
// выводя их путем изменения третьего аргумента функции SetColorlndex некоторым нетривиальным способом.
			colorIdx :=uint8(rand.Intn(len(palette) - 1) + 1)
			img.SetColorIndex(size + int(x*size+0.5), size + int(y*size+0.5), colorIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
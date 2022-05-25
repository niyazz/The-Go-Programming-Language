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
	"strconv"
	"os"
)

var green = color.RGBA{0x12, 0x80, 0x12, 0xff}
var red = color.RGBA{0x90, 0x10, 0x12, 0xff}
var blue = color.RGBA{0x12, 0x10, 0x90, 0xff}
var palette = []color.Color{color.Black, green, red, blue}
var mu sync.Mutex
var count int

func main(){
	handler := func(w http.ResponseWriter, r *http.Request){

// Task 1.12: go run main.go
// Измените сервер с фигурами Лиссажу так, чтобы значения 
// параметров считывались из URL. Например, URL вида http://localhost:8000/ 
// ?cycles=20 устанавливает количество циклов равным 20 вместо значения по умолчанию, равного 5
		q := r.URL.Query()
		if(len(q["cycles"]) >= 1){
			cycles, err := strconv.Atoi(q["cycles"][0])
			if err != nil{
				fmt.Printf(err.Error())
				os.Exit(1)
			}
			lissajous(w, cycles)
		}
		lissajous(w, 5)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int){
	const(
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIdx :=uint8(rand.Intn(len(palette) - 1) + 1)
			img.SetColorIndex(size + int(x*size+0.5), size + int(y*size+0.5), colorIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
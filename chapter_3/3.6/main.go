package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height = 1024, 1024
	)

	f, _ := os.Create("task36.png")
	img:= image.NewRGBA(image.Rect(0, 0, width, height))
	for py:=0; py < height; py++{
		y1 := float64(py) / height * (ymax - ymin) + ymin
		y2 := float64(py + 1) / height * (ymax - ymin) + ymin
		for px:=0;px <width;px++{
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px + 1)/width*(xmax-xmin) + xmin

			img.Set(px, py, getColor(x1, x2, y1, y2))
		}
	}
	png.Encode(f, img)
	f.Close()
}

// Returns color of square of 4 side length 
func getColor(x1, x2, y1, y2 float64) color.Color{
	m1 := mandelbrot(complex(x1, y1))
	m2 := mandelbrot(complex(x2, y1))
	m3 := mandelbrot(complex(x1, y2))
	m4 := mandelbrot(complex(x2, y2))

	r1, g1, b1, a1 :=m1.RGBA()
	r2, g2, b2, a2 :=m2.RGBA()
	r3, g3, b3, a3 :=m3.RGBA()
	r4, g4, b4, a4 :=m4.RGBA()

	var size uint32 = 4

	r := (r1 + r2 + r3 +r4)/size
	g := (g1 + g2 + g3 +g4)/size
	b := (b1 + b2 + b3 +b4)/size
	a := (a1 + a2 + a3 +a4)/size

	return color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++{
		v = v*v + z
		if cmplx.Abs(v) > 2{
			return color.RGBA{255- contrast*n, 0, 0, 255}
		}
	}
	return color.RGBA{0, 100, 0, 255}
}
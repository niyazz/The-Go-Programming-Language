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

	f, _ := os.Create("task35.png")
	img:= image.NewRGBA(image.Rect(0, 0, width, height))
	for py:=0; py < height; py++{
		y := float64(py) / height * (ymax - ymin) + ymin
		for px:=0;px <width;px++{
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, newton(z))
		}
	}
	png.Encode(f, img)
	f.Close()
}

//https://en.wikipedia.org/wiki/Newton_fractal#:~:text=The%20Newton%20fractal%20is%20a,is%20given%20by%20Newton's%20method.
// f(z) = z^4 - 1
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color{
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++{
		z = z - (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{255- contrast*n, 0, 0, 255}
		}
	}
	return color.RGBA{0, 100, 0, 255}
}
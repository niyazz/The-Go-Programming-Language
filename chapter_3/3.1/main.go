package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	xyscale = width/2/xyrange
	zscale = height * 0.4
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main(){
	f, _ := os.Create("task31.txt")
	s := fmt.Sprintf("<svg xmlns='http://w3.org/2000/svg' "+
"style='stroke: grey; fill:white; stroke-width: 0.7' " +
"width='%d' height='%d'>", width, height)
	fmt.Fprint(f, s)

	for i:=0; i < cells; i++{
		for j:=0; j < cells;j++{
			ax,ay,e := corner(i+1, j)

			if e != nil {
				continue
			}
			bx,by,e := corner(i, j)
			if e != nil {
				continue
			}
			cx,cy,e := corner(i, j+1)
			if e != nil {
				continue
			}
			dx,dy,e := corner(i+1, j+1)
			if e != nil {
				continue
			}
	
			s := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
		ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Fprintf(f, s)
		}
	}
	fmt.Fprintf(f, "</svg>")
	f.Close()
}

func corner(i, j int) (float64, float64, error){
	x:= xyrange * (float64(i)/cells - 0.5)
	y:= xyrange * (float64(j)/cells - 0.5)
	z := f(x,y)

	// Task 3.1: go run .
	// Если функция f возвращает значение float64, не являющееся конечным, SVG-файл содержит неверные элементы <polygon>.
    // Измените программу так, чтобы некорректные многоугольники были опущены.
	if math.IsInf(z, 0) || math.IsNaN(z){
		return 0, 0, fmt.Errorf("invalid value")
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64{
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

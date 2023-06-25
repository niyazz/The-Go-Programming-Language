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
	f, _ := os.Create("task33.svg")
	s := fmt.Sprintf("<svg xmlns='http://w3.org/2000/svg' "+
"style='stroke: grey; fill:white; stroke-width: 0.7' " +
"width='%d' height='%d'>", width, height)
	fmt.Fprint(f, s)

	for i:=0; i < cells; i++{
		for j:=0; j < cells;j++{
			ax,ay,isPeak1,e := corner(i+1, j)
			if e != nil {
				continue
			}
			bx,by,isPeak2,e := corner(i, j)
			if e != nil {
				continue
			}
			cx,cy,isPeak3,e := corner(i, j+1)
			if e != nil {
				continue
			}
	
			dx,dy,isPeak4,e := corner(i+1, j+1)
			if e != nil {
				continue
			}
			color:=getColor(isPeak1 && isPeak2 && isPeak3 && isPeak4);
	
			s := fmt.Sprintf("<polygon style ='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
		color, ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Fprintf(f, s)
		}
	}
	fmt.Fprintf(f, "</svg>")
	f.Close()
}


// Task 3.3: go run .
// Окрасьте каждый многоугольник цветом, зависящим от его высоты, 
// так, чтобы пики были красными (#ff0000), а низины — синими (#0000ff).
func getColor(b bool) string{
	if b{
		return "red"
	}
	return "blue"
}

func corner(i, j int) (float64, float64, bool, error){
	x:= xyrange * (float64(i)/cells - 0.5)
	y:= xyrange * (float64(j)/cells - 0.5)
	z := f(x,y)


	if math.IsInf(z, 0) || math.IsNaN(z){
		return 0, 0, false, fmt.Errorf("invalid value")
	}

	isPeak := z >= 0

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isPeak, nil
}

func f(x, y float64) float64{
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggsBox(x, y float64) float64{
	a := 3.0
	b := 25.0
	return a * (math.Sin(x)/b + math.Sin(y)/b)
}
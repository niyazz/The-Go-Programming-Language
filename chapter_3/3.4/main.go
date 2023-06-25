package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	angle = math.Pi / 6
	defmincolor = "blue"
	defmaxcolor = "red"
)

var xyscale float64
var zscale float64
var sin30, cos30 = math.Sin(angle), math.Cos(angle)


// Task 3.4: go run .
// url: http://localhost:8000/?minc=green&w=300&h=600&maxc=red
// note: иногда exe может вовсе не заупскать из-за антивирусника, подробнее https://go.dev/doc/faq#virus
// Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7, создайте веб-сервер, который вычисляет поверхности и возвращает
// клиенту SVG-данные. Сервер должен использовать в ответе заголовок ContentType наподобие следующего:
func main(){
	var svg string
	handler := func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("ContentType", "image/svg+xml")
	wid, h, minc, maxc := parseParams(r.URL.Query())
	
	if wid != 0 && h != 0 && minc !="" && maxc != ""{
		svg = getSVG(wid, h, maxc, minc)
	}else{
		svg = getSVG(width, height, defmaxcolor, defmincolor)
	}

	io.WriteString(w, svg)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseParams(params url.Values) (int, int, string, string){
	var width, height int
	var maxcolor, mincolor string
	if v, ok := params["w"]; ok{
		width, _ = strconv.Atoi(v[0])
	}
	if v, ok := params["h"]; ok{
		height, _ = strconv.Atoi(v[0])
	}
	if v, ok := params["maxc"]; ok{
		maxcolor = v[0]
	}
	if v, ok := params["minc"]; ok{
		mincolor = v[0]
	}
	return width, height, mincolor, maxcolor
}

func getSVG(w, h int, maxcolor, mincolor string) string{
	xyscale = float64(w/2/xyrange)
	zscale = float64(h) * 0.4
	SVGcode := fmt.Sprintf("<svg xmlns='http://w3.org/2000/svg' "+
	"style='stroke: grey; fill:white; stroke-width: 0.7' " +
	"width='%d' height='%d'>", w, h)

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
			color:=getColor(isPeak1 && isPeak2 && isPeak3 && isPeak4, mincolor, maxcolor);
	
			SVGcode += fmt.Sprintf("<polygon style ='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
		color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	SVGcode +=  "</svg>"
	return SVGcode
}

func getColor(b bool, c1, c2 string) string{
	if b{
		return c2
	}
	return c1
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
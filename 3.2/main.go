package main


import (
	"fmt"
	"math"
	"io"
	"net/http"
	"log"
)

const (
	width, height 	= 600, 320
	cells 			= 100
	xyrange			= 30.0
	xyscale			= width / 2 / xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var zmax, zmin float64
func init() {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := z(i, j)
			if z < zmin {
				zmin = z
			}
			if (zmax < z) {
				zmax = z
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}
func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg	xmlns='http://www.w3.org/2000/svg'	"+
		"style='stroke:	grey;	fill:	white;	stroke-width:	0.7'	"+
		"width='%d'	height='%d'>",	width,	height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1, j + 1)
			color := color(i, j)
			fmt.Fprintf(out, "<polygon	points='%g,%g	%g,%g	%g,%g	%g,%g' style=\"fill: #%x;\"/>\n",
				ax,	ay,	bx,	by,	cx,	cy,	dx,	dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}
func color(i, j int) uint32 {
	z := z(i, j)
	z = (z - zmin) / (zmax - zmin)
	color := uint32((0xff0000 * z)) & 0xff0000 + uint32((1 - z) * 0x0000ff) & 0x0000ff
	return color
}
func z(i, j int) float64 {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	z := f(x, y)
	return z
}
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	z := f(x, y)

	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
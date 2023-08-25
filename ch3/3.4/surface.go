// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320              // canvas size in pixels
	xyscale       = width / 2 / xyrange   // pixels per x or y unit
	zscale        = float64(height) * 0.4 // pixels per z unit
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
	red     = "#ff0000"
	blue    = "#0000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).

	sx := float64(width)/2 + (x-y)*cos30*float64(xyscale)
	sy := float64(height)/2 + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy, z >= 0
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// IsInf 第二个参数 > 0 判断是否正无穷， < 0 判断是否负无穷， = 0 判断是否无穷
func isInfinite(x float64) bool {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return true
	}
	return false
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if r.Form["width"] != nil {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if r.Form["height"] != nil {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, isPeak := corner(i+1, j+1)

			if isInfinite(ax) || isInfinite(ay) || isInfinite(bx) || isInfinite(by) || isInfinite(cx) || isInfinite(cy) || isInfinite(dx) || isInfinite(dy) {
				continue
			}

			color := blue
			if isPeak {
				color = red
			}

			fmt.Fprintf(w, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

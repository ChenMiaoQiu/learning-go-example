// Surface computes an SVG rendering of a 3-D surface function.
// go run surface.go > test.svg
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if isInfinite(ax) || isInfinite(ay) || isInfinite(bx) || isInfinite(by) || isInfinite(cx) || isInfinite(cy) || isInfinite(dx) || isInfinite(dy) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := eggBox(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	if isInfinite(r) {
		return 0
	}

	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 10
}

func saddle(x, y float64) float64 {
	return math.Pow(y/xyrange*2, 2) - math.Pow(x/xyrange*2, 2)
}

func moguls(x, y float64) float64 {
	return math.Pow(math.Sin(x/xyrange*3*math.Pi), 2) * math.Cos(y/xyrange*3*math.Pi)
}

// IsInf 第二个参数 > 0 判断是否正无穷， < 0 判断是否负无穷， = 0 判断是否无穷
func isInfinite(x float64) bool {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return true
	}
	return false
}

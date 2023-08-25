// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	x, y, zoom := 0.0, 0.0, 0.0

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if r.Form["x"] != nil {
		x, _ = strconv.ParseFloat(r.Form["x"][0], 64)
	}

	if r.Form["y"] != nil {
		y, _ = strconv.ParseFloat(r.Form["y"][0], 64)
	}

	if r.Form["zoom"] != nil {
		zoom, _ = strconv.ParseFloat(r.Form["zoom"][0], 64)
	}
	render(w, x, y, zoom)
}

func render(out io.Writer, x, y, zoom float64) {
	const (
		// xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)
	exp := math.Exp2(1 - zoom)
	xmin, xmax := x-exp, x+exp
	ymin, ymax := y-exp, y+exp

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.White, color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff}, color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0x22, 0x22, 0x22, 0xff}, color.RGBA{0xff, 0x44, 0x15, 0xff},
	color.RGBA{0x44, 0x62, 0x12, 0xff}, color.RGBA{0xff, 0xff, 0x15, 0xff},
}

func lissajous(cycles int, out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(7)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
/*
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
*/

// handler echoes the HTTP request.

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	var cycle string
	for k, v := range r.Form {
		if k == "cycle" {
			cycle = v[0]
		}
	}

	cycle_int, err := strconv.Atoi(cycle)
	if err != nil {
		cycle_int = 5
	}

	lissajous(cycle_int, w)

	mu.Unlock()
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

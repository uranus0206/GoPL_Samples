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

func main() {
	http.HandleFunc("/v2/", handler)
	http.HandleFunc("/v3/", handlerV2)
	http.HandleFunc("/count", counter)

	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		cycles := 0
		for k, v := range r.Form {
			if k == "cycles" {
				cycles, _ = strconv.Atoi(v[0])
			}
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

// handler echos the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// server3 handler echos the HTTP request.
func handlerV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v) // %q for show quoted of string
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echos the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

var colorGreen = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var colorRed = color.RGBA{0xff, 0x00, 0x00, 0xFF}
var colorBlue = color.RGBA{0x00, 0x00, 0xff, 0xFF}

var palette = []color.Color{color.White, color.Black, colorGreen, colorBlue, colorRed} // Instantiating a slice

// Golbal const
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette

	greenIndex = 2
	blueIndex  = 3
	redIndex   = 4
)

func lissajous(out io.Writer, cycle int) {
	// Local const
	const (
		// cycles  =  5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	cycles := 5
	if cycle > 0 {
		cycles = cycle
	}

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // Instantiate a struct
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		lineColorIndex := rand.Intn(3) + 1
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// img.SetColorIndex(size+int(x*size+0.5),
			// 	size+int(y*size+0.5),
			// 	blackIndex)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				uint8(lineColorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

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

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	// Local const
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // Instantiate a struct
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		lineColorIndex := rand.Intn(3) + 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
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

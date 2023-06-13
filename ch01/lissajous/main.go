// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff}}

// the value of a constant must be a number, string, or boolean
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2 // next color in palette
)

func main() {
	file, err := os.Create("lissajous.gif")
	if err != nil {
		log.Fatal(err)
	}
	//lissajous(os.Stdout)
	lissajous(file)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5     // number of complete x oscillator revolutions
		res    = 0.001 // angular resolution
		size   = 200   // image canvas covers [-size..+size]
		frames = 64    // number of animation frames
		delay  = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 5.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: frames}
	phase := 0.0 // phase difference
	paletteColorIndex := blackIndex
	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			if paletteColorIndex == blackIndex {
				paletteColorIndex = greenIndex
			} else {
				paletteColorIndex = blackIndex
			}
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(paletteColorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}
}

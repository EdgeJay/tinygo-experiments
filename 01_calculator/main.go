package main

import (
	"image/color"
	"time"

	"github.com/edgejay/tinygo-experiments/internal/display/ssd1306"
	"github.com/edgejay/tinygo-experiments/internal/machine/rp2040"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	rp2040.ConfigureMachine()
	display := ssd1306.ConfigureDisplay(true)
	display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)

	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}

	// tinyfont.WriteLine(display, &freemono.Bold9pt7b, 5, 12, "CALCULATOR", white)

	// create int var that stores ASCII value of A to Z (65 to 90) characters
	asciiVal := 65

	for {
		display.ClearDisplay()
		tinyfont.WriteLine(display, &gophers.Regular32pt, 5, 50, string(rune(asciiVal)), white)
		asciiVal++
		if asciiVal > 90 {
			asciiVal = 65
		}
		display.Display()
		time.Sleep(1000 * time.Millisecond)
	}
}

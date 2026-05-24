package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 2.8 * machine.MHz,
		SDA:       machine.GPIO12,
		SCL:       machine.GPIO13,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address:  0x3C,
		Width:    128,
		Height:   64,
		Rotation: drivers.Rotation180,
	})
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

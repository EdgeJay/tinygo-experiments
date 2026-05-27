package display

import (
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/gophers"
)

// ShowGophers is for testing the gophers font from tinyfont. It will display the ASCII characters
// from A to Z on the display, one at a time, with a delay of 1 second between each character.
// Take note this function is a blocking function and will run indefinitely until the device is reset or powered off.
func ShowGophers(display *ssd1306.Device) {
	// create int var that stores ASCII value of A to Z (65 to 90) characters
	asciiVal := 65

	for {
		display.ClearDisplay()
		tinyfont.WriteLine(display, &gophers.Regular32pt, 5, 50, string(rune(asciiVal)), whiteColour)
		asciiVal++
		if asciiVal > 90 {
			asciiVal = 65
		}
		display.Display()
		time.Sleep(1000 * time.Millisecond)
	}
}

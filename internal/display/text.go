package display

import (
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func ShowText(display *ssd1306.Device, x, y int16, text string) {
	tinyfont.WriteLine(display, &freemono.Bold9pt7b, x, y, text, whiteColour)
	display.Display()
}

package main

import (
	"time"

	disp "github.com/edgejay/tinygo-experiments/internal/display"
	"github.com/edgejay/tinygo-experiments/internal/display/ssd1306"
	"github.com/edgejay/tinygo-experiments/internal/machine/rp2040"
)

func main() {
	rp2040.ConfigureMachine()
	display := ssd1306.ConfigureDisplay(true)
	display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)

	disp.ShowText(display, 5, 12, "CALCULATOR")
}

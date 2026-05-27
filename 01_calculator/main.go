package main

import (
	"time"

	disp "github.com/edgejay/tinygo-experiments/internal/display"
	"github.com/edgejay/tinygo-experiments/internal/display/ssd1306"
	"github.com/edgejay/tinygo-experiments/internal/keyboard"
	"github.com/edgejay/tinygo-experiments/internal/machine/rp2040"
)

func main() {
	rp2040.ConfigureMachine()
	display := ssd1306.ConfigureDisplay(true)
	display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)

	disp.ShowText(display, 5, 12, "CALCULATOR")

	kb, err := keyboard.NewKeyboard(map[int]rune{
		0:  '7',
		1:  '4',
		2:  '1',
		3:  '8',
		4:  '5',
		5:  '2',
		6:  '9',
		7:  '6',
		8:  '3',
		9:  '0',
		10: '.',
		11: '=',
	})

	if err != nil {
		panic(err)
	}

	keyCh := make(chan rune)

	go func() {
		kb.Listen(keyCh)
	}()

	displayedText := ""

	for {
		select {
		case key := <-keyCh:
			display.ClearDisplay()
			displayedText += string(key)
			disp.ShowText(display, 5, 30, displayedText)
		}
	}
}

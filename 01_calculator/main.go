package main

import (
	"time"

	disp "github.com/edgejay/tinygo-experiments/internal/display"
	dd "github.com/edgejay/tinygo-experiments/internal/display/ssd1306"
	"github.com/edgejay/tinygo-experiments/internal/joystick"
	jsm "github.com/edgejay/tinygo-experiments/internal/joystick/mapping"
	"github.com/edgejay/tinygo-experiments/internal/keyboard"
	kbm "github.com/edgejay/tinygo-experiments/internal/keyboard/mapping"
	"github.com/edgejay/tinygo-experiments/internal/machine/rp2040"
	"tinygo.org/x/drivers/ssd1306"
)

func setupDisplay() *ssd1306.Device {
	display := dd.ConfigureDisplay(true)
	display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)
	disp.ShowText(display, 5, 12, "CALCULATOR")
	return display
}

func setupKeyboard() (*keyboard.Keyboard, chan rune, error) {
	kb, err := keyboard.NewKeyboard(kbm.GetCalculatorKeysMapping())
	if err != nil {
		return nil, nil, err
	}

	keyCh := make(chan rune)

	go func() {
		kb.Listen(keyCh)
	}()

	return kb, keyCh, nil
}

func setupJoystick() (*joystick.Joystick, chan joystick.JoystickState) {
	js := joystick.NewJoystick(jsm.GetCalculatorKeysMapping())
	jsCh := make(chan joystick.JoystickState)

	go func() {
		js.Listen(jsCh, 250*time.Millisecond)
	}()

	return js, jsCh
}

func main() {
	rp2040.ConfigureMachine()
	display := setupDisplay()

	_, keyCh, err := setupKeyboard()
	if err != nil {
		panic(err)
	}

	js, jsCh := setupJoystick()

	displayedText := ""
	for {
		select {
		case key := <-keyCh:
			display.ClearDisplay()
			displayedText += string(key)
			disp.ShowText(display, 5, 30, displayedText)
		case jsState := <-jsCh:
			if jsState.IsNeutral() {
				continue
			}

			display.ClearDisplay()
			if jsState.CenterButtonPressed {
				displayedText = ""
			}
			if jsState.Up {
				displayedText += js.GetKey("up")
			}
			if jsState.Left {
				displayedText += js.GetKey("left")
			}
			if jsState.Down {
				displayedText += js.GetKey("down")
			}
			if jsState.Right {
				displayedText += js.GetKey("right")
			}
			disp.ShowText(display, 5, 30, displayedText)
		}
	}
}

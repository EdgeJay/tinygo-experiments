package main

import (
	"strconv"
	"time"

	"github.com/edgejay/tinygo-experiments/internal/calculator"
	disp "github.com/edgejay/tinygo-experiments/internal/display"
	dd "github.com/edgejay/tinygo-experiments/internal/display/ssd1306"
	"github.com/edgejay/tinygo-experiments/internal/joystick"
	jsm "github.com/edgejay/tinygo-experiments/internal/joystick/mapping"
	"github.com/edgejay/tinygo-experiments/internal/keyboard"
	kbm "github.com/edgejay/tinygo-experiments/internal/keyboard/mapping"
	"github.com/edgejay/tinygo-experiments/internal/machine/rp2040"
	"tinygo.org/x/drivers/ssd1306"
)

const (
	displayTextX        = 5
	displayTextY        = 42
	displayTextLenLimit = 10
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
		js.Listen(jsCh, 200*time.Millisecond)
	}()

	return js, jsCh
}

func isLastCharOperator(s string) bool {
	if len(s) == 0 {
		return false
	}

	lastChar := s[len(s)-1]
	return lastChar == '+' || lastChar == '-' || lastChar == '*' || lastChar == '/'
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
			if key != '=' {
				displayedText += string(key)
			} else {
				// do calculation
				calc := calculator.NewCalculator(displayedText)
				result, err := calc.Calculate()
				if err != nil {
					displayedText = "Error"
				} else {
					displayedText = strconv.FormatFloat(float64(result), 'f', -1, 64)
				}
			}

			// update display
			disp.ShowText(display, displayTextX, displayTextY, displayedText)
		case jsState := <-jsCh:
			// clear display if center button is pressed
			if jsState.CenterButtonPressed {
				display.ClearDisplay()
				displayedText = ""
				disp.ShowText(display, displayTextX, displayTextY, displayedText)
				continue
			}

			// skip if:
			// joystick is in the neutral position, or
			// last character is an operator, or
			// displayed text is empty (to prevent starting with an operator)
			if jsState.IsNeutral() || isLastCharOperator(displayedText) || displayedText == "" {
				continue
			}

			display.ClearDisplay()

			// add math operators based on the joystick direction
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

			// update display
			disp.ShowText(display, displayTextX, displayTextY, displayedText)
		}
	}
}

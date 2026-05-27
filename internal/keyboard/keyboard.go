package keyboard

import (
	"errors"
	"machine"
	"time"
)

var colPins = []machine.Pin{
	machine.GPIO5,
	machine.GPIO6,
	machine.GPIO7,
	machine.GPIO8,
}

var rowPins = []machine.Pin{
	machine.GPIO9,
	machine.GPIO10,
	machine.GPIO11,
}

var ErrInvalidKeysMapping = errors.New("invalid keys mapping")

type Keyboard struct {
	keysMapping map[int]rune
}

func init() {
	println("Initializing keyboard...")

	// init pin config
	for _, c := range colPins {
		c.Configure(machine.PinConfig{Mode: machine.PinOutput})
		c.Low()
	}

	for _, c := range rowPins {
		c.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}
}

func resetColPins() {
	for _, c := range colPins {
		c.Low()
	}
}

func NewKeyboard(keysMapping map[int]rune) (*Keyboard, error) {
	// keysMapping must be a map of 12 keys (0 to 11) to their corresponding rune values, for example:
	// keysMapping := map[int]rune{
	// 	0: 'A',
	// 	1: 'B',
	// ...
	// 	11: 'C',
	// }
	if len(keysMapping) != 12 {
		return nil, ErrInvalidKeysMapping
	}

	return &Keyboard{keysMapping}, nil
}

func (kb *Keyboard) Listen(cb chan<- rune) {
	for {
		for colIdx, c := range colPins {
			resetColPins()
			c.High()
			time.Sleep(1 * time.Millisecond)

			for rowIdx, r := range rowPins {
				if r.Get() { // key is pressed
					keyIdx := rowIdx + colIdx*len(rowPins)
					cb <- kb.keysMapping[keyIdx]
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}
}

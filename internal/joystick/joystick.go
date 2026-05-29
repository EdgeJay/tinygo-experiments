package joystick

import (
	"machine"
	"time"

	"github.com/edgejay/tinygo-experiments/internal/joystick/mapping"
)

type Joystick struct {
	X               machine.ADC
	Y               machine.ADC
	CenterButton    machine.Pin
	JoystickMapping mapping.JoystickMapping
}

type JoystickState struct {
	Up                  bool
	Down                bool
	Left                bool
	Right               bool
	CenterButtonPressed bool
}

func (jss JoystickState) IsNeutral() bool {
	return !jss.Up && !jss.Down && !jss.Left && !jss.Right && !jss.CenterButtonPressed
}

func init() {
	println("Initializing joystick...")
	machine.InitADC()
}

func NewJoystick(jsm mapping.JoystickMapping) *Joystick {
	joystickX := machine.ADC{Pin: machine.GPIO29}
	joystickX.Configure(machine.ADCConfig{})
	joystickY := machine.ADC{Pin: machine.GPIO28}
	joystickY.Configure(machine.ADCConfig{})

	joystickButton := machine.GPIO0
	joystickButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return &Joystick{
		X:               joystickX,
		Y:               joystickY,
		CenterButton:    joystickButton,
		JoystickMapping: jsm,
	}
}

func (js *Joystick) Listen(cb chan<- JoystickState, delay time.Duration) {
	for {
		xValue := js.X.Get()
		yValue := js.Y.Get()

		var state JoystickState
		state.Up = yValue > 0xA000
		state.Down = yValue < 0x6000
		state.Left = xValue < 0x6000
		state.Right = xValue > 0xA000
		state.CenterButtonPressed = !js.CenterButton.Get() // active low

		// send state to channel
		cb <- state

		time.Sleep(delay) // adjust as needed for responsiveness
	}
}

func (js *Joystick) GetKey(joystickKey string) string {
	switch joystickKey {
	case "up":
		return js.JoystickMapping.Up
	case "down":
		return js.JoystickMapping.Down
	case "left":
		return js.JoystickMapping.Left
	case "right":
		return js.JoystickMapping.Right
	case "center":
		return js.JoystickMapping.Center
	default:
		return ""
	}
}

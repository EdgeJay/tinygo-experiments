package joystick

import (
	"machine"
	"time"
)

type Joystick struct {
	X            machine.ADC
	Y            machine.ADC
	CenterButton machine.Pin
}

type JoystickState struct {
	Up                  bool
	Down                bool
	Left                bool
	Right               bool
	CenterButtonPressed bool
}

func init() {
	println("Initializing joystick...")
	machine.InitADC()
}

func NewJoystick() *Joystick {
	joystickX := machine.ADC{Pin: machine.GPIO29}
	joystickX.Configure(machine.ADCConfig{})
	joystickY := machine.ADC{Pin: machine.GPIO28}
	joystickY.Configure(machine.ADCConfig{})

	joystickButton := machine.GPIO0
	joystickButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return &Joystick{
		X:            joystickX,
		Y:            joystickY,
		CenterButton: joystickButton,
	}
}

func (js *Joystick) Listen(cb chan<- JoystickState) {
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

		time.Sleep(100 * time.Millisecond) // adjust as needed for responsiveness
	}
}

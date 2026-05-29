package rotary

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/encoders"
)

var rotaryEncoder *encoders.QuadratureDevice

func init() {
	rotaryEncoder = encoders.NewQuadratureViaInterrupt(
		machine.GPIO4,
		machine.GPIO3,
	)
}

type RotaryState struct {
	Left  bool
	Right bool
	Value int
}

type RotaryEncoder struct {
	Value          int
	MinValue       int
	MaxValue       int
	oldRotaryValue int
}

func NewRotaryEncoder(initValue, minValue, maxValue, precision int) *RotaryEncoder {
	rotaryEncoder.Configure(encoders.QuadratureConfig{
		Precision: precision,
	})

	return &RotaryEncoder{
		Value:          initValue,
		MinValue:       minValue,
		MaxValue:       maxValue,
		oldRotaryValue: 0,
	}
}

func (re *RotaryEncoder) Listen(cb chan<- RotaryState) {
	for {
		state := RotaryState{
			Left:  false,
			Right: false,
			Value: re.Value,
		}

		if newValue := rotaryEncoder.Position(); newValue != re.oldRotaryValue {
			if newValue < re.oldRotaryValue {
				state.Right = true
				state.Value = min(re.Value+1, re.MaxValue)
			} else {
				state.Left = true
				state.Value = max(re.Value-1, re.MinValue)
			}
			re.Value = state.Value
			re.oldRotaryValue = newValue
			cb <- state
		}

		time.Sleep(100 * time.Millisecond)
	}
}

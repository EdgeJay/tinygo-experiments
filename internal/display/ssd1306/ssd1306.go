package ssd1306

import (
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
)

func ConfigureDisplay(rotate180 bool) *ssd1306.Device {
	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
		Rotation: func() drivers.Rotation {
			if rotate180 {
				return drivers.Rotation180
			}
			return drivers.Rotation0
		}(),
	})
	return display
}

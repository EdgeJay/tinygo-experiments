package led

import (
	"machine"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
)

type WS2812B struct {
	ws *piolib.WS2812B
}

func NewWS2812B() (*WS2812B, error) {
	s, err := pio.PIO0.ClaimStateMachine()
	if err != nil {
		return nil, err
	}

	ws, err := piolib.NewWS2812B(s, machine.GPIO1)
	if err != nil {
		return nil, err
	}

	ws.EnableDMA(true)
	return &WS2812B{
		ws: ws,
	}, nil
}

func (ws *WS2812B) GetAllOffColours() []uint32 {
	return []uint32{
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
	}
}

func (ws *WS2812B) PutColour(ledPosition int, colour uint32) error {
	colours := ws.GetAllOffColours()
	colours[ledPosition] = colour
	return ws.WriteRaw(colours)
}

func (ws *WS2812B) WriteRaw(rawGRB []uint32) error {
	return ws.ws.WriteRaw(rawGRB)
}

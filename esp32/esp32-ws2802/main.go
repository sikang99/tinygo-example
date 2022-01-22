package main

// This is an example of controlling WS2812 LEDs from an ESP32.
// The following PRs are still needed to get this to work:
// https://github.com/tinygo-org/tinygo/pull/1353
// https://github.com/tinygo-org/tinygo/pull/1354
// https://github.com/tinygo-org/drivers/pull/198

import (
	"machine"
	"time"

	"github.com/aykevl/ledsgo"
	"tinygo.org/x/drivers/ws2812"
)

const brightness = 0x44
const pin = machine.Pin(27)

var ws ws2812.Device

func main() {
	strip := make(ledsgo.Strip, 50)
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = ws2812.New(pin)
	rainbow(strip)
}

func rainbow(strip ledsgo.Strip) {
	for {
		now := time.Now().UnixNano()
		for i := range strip {
			strip[i] = ledsgo.Color{uint16(now>>15) - uint16(i)<<12, 0xff, brightness}.Spectrum()
		}
		ws.WriteColors(strip)
		time.Sleep(time.Second / 100)
	}
}

func noise(strip ledsgo.Strip) {
	for {
		now := time.Now().UnixNano()
		for i := range strip {
			const spread = 100
			val := int32(ledsgo.Noise2(int32(now>>22), int32(i*spread))) * 3 / 2
			strip[i] = ledsgo.Color{uint16(val), 0xff, brightness}.Spectrum()
		}
		ws.WriteColors(strip)
		time.Sleep(time.Second / 100)
	}
}

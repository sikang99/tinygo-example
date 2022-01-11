package main

import (
	"machine"
	"time"
)

// start here at main function
func main() {
	go blink(machine.LED1, 1000*time.Millisecond)
	go blink(machine.LED2, 750*time.Millisecond)
}

// blink the LED with given duration
func blink(led machine.Pin, delay time.Duration) {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(delay)
		led.High()
		time.Sleep(delay)
	}
}

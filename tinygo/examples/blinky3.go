package main

import (
	"machine"
	"time"
)

func led_light(changeChan chan bool, led machine.Pin) {
	for {
		led.Set(<-changeChan)
	}
}

const (
	buttonMode      = machine.PinInputPulldown
	buttonPinChange = machine.PinRising | machine.PinFalling
)

type ButtonDelegate struct {
	changeChan chan bool
}

func NewButtonDelegate(changeChan chan bool) *ButtonDelegate {
	return &ButtonDelegate{changeChan: changeChan}
}
func (b *ButtonDelegate) buttonChange(pin machine.Pin) {
	b.changeChan <- pin.Get()
}

func busyLedBlink(led machine.Pin) {
	for {
		led.Low()
		time.Sleep(time.Millisecond * 100)

		led.High()
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	buttonChangeChan := make(chan bool, 10)

	led1 := machine.LED1
	led1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.BUTTON
	button.Configure(machine.PinConfig{Mode: buttonMode})

	button.SetInterrupt(buttonPinChange, NewButtonDelegate(buttonChangeChan).buttonChange)

	led2 := machine.LED2
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go led_light(buttonChangeChan, led1)
	busyLedBlink(led2)

}

package controller

import (
	"fmt"
	"os"
)

type pwmPin struct {
	//GPIO Pin Number
	Pin   string
	Value float32
}

const ON = 1
const OFF = 0

func newPWMPin(pinNumber string) (pin pwmPin, err error) {

	pin.Pin = pinNumber
	pin.Set(OFF)
	return pin, nil

}

func (p *pwmPin) Set(value float32) error {

	if value < OFF {
		value = OFF
	} else if value > ON {
		value = ON
	}

	data := fmt.Sprintf("%s=%f\n", p.Pin, value)

	f, err := os.OpenFile("/dev/pi-blaster", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}

	f.Sync()

	p.Value = value

	return nil
}

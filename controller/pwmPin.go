package controller

import (
	"fmt"
	"os"
)

type pwmPin struct {
	//GPIO Pin Number
	Pin   string
	Value float64
}

const ON = 1
const OFF = 0

func newPWMPin(pinNumber string) (*pwmPin, error) {

	p := pwmPin{Pin: pinNumber}

	err := p.Set(OFF)
	if err != nil {
		return nil, err
	}
	return &p, nil

}

func (p *pwmPin) Set(value float64) error {

	value = clamp(value, OFF, ON)

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

	p.Value = value

	return nil
}

package controller

import (
	"fmt"
	"os"
)

type pwmPin struct {
	//GPIO Pin Number
	Pin   string
	Value byte
}

func newPWMPin(pinNumber string) (pin pwmPin, err error) {

	pin.Pin = pinNumber
	pin.SetAnalog(0)
	return pin, nil

}

func (p *pwmPin) SetAnalog(value byte) error {
	percentageValue := float32(value) / 255

	data := fmt.Sprintf("%s=%f\n", p.Pin, percentageValue)

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

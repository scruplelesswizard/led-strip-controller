package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type StripDef struct {
	Name     string `yaml:"name"`
	RedPin   string `yaml:"redPin"`
	GreenPin string `yaml:"greenPin"`
	BluePin  string `yaml:"bluePin"`
}

type StripsDef struct {
	Strips []*StripDef `yaml:"strips"`
}

func LoadStripsDefFromFile(path string) (*StripsDef, error) {

	var sD *StripsDef

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Could not stat %s: %s", path, err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Unable to read %s: %s", path, err)
		return sD, err
	}

	err = yaml.Unmarshal(data, &sD)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshal %s to YAML: %s", path, err)
		return sD, err
	}

	return sD, nil

}

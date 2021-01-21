package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/tarm/serial"
	"gopkg.in/yaml.v2"
)

const ConfPath = "./config.yml"

type Conf struct {
	Serial SerialConf `yaml:"serial"`
	Sender SenderConf `yaml:"sender"`
}

func NewConf() *Conf {
	c := readConf(true)
	return c
}

func readConf(verbose bool) *Conf {
	yamlFile, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c Conf
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}

	if verbose {
		fmt.Printf("configuration parameter: %+v\n", c)
	}

	return &c
}

func (c *Conf) GetSerialConf() *serial.Config {
	s := c.Serial
	sc := &serial.Config{
		Name:     s.Name,
		Baud:     s.Baud,
		Size:     getSize(s),
		Parity:   getParity(s),
		StopBits: getStopbits(s),
		// NOTE: Setting `ReadTimeout` will not block on goroutine. Keep comment out.
		ReadTimeout: time.Millisecond * 1000,
	}
	return sc
}

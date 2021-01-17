package myserial

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tarm/serial"
	"gopkg.in/yaml.v2"
	// "time"
)

type conf struct {
	Name     string `yaml:"name"`
	Baud     int    `yaml:"baud"`
	Databits int    `yaml:"databits"`
	Parity   string `yaml:"parity"`
	Stopbits int    `yaml:"stopbits"`
}

func (c *conf) readConf(path string) *conf {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *conf) getSize() byte {
	return byte(c.Databits)
}

func (c *conf) getParity() serial.Parity {
	switch c.Parity {
	case "None":
		return serial.ParityNone
	default:
		log.Printf("default parity selected: None")
		return serial.ParityNone
	}
}

func (c *conf) getStopbits() serial.StopBits {
	switch c.Stopbits {
	case 1:
		return serial.Stop1
	case 15:
		return serial.Stop1Half
	case 2:
		return serial.Stop2
	default:
		log.Printf("default stopbits selected: 1")
		return serial.Stop1
	}
}

func GetConf(path string, verbose bool) *serial.Config {
	var c conf
	c.readConf(path)
	if verbose {
		fmt.Printf("configuration parameter: %+v\n", c)
	}

	sc := &serial.Config{
		Name:     c.Name,
		Baud:     c.Baud,
		Size:     c.getSize(),
		Parity:   c.getParity(),
		StopBits: c.getStopbits(),
		// NOTE: Setting `ReadTimeout` will not block on goroutine. Keep comment out.
		// ReadTimeout: time.Millisecond * 5000,
	}
	return sc
}

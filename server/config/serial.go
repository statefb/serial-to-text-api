package config

import (
	"log"

	"github.com/tarm/serial"
)

type SerialConf struct {
	Mock         bool   `yaml:"mock"`
	EndSignature string `yaml:"endSignature"`
	Name         string `yaml:"name"`
	Baud         int    `yaml:"baud"`
	Databits     int    `yaml:"databits"`
	Parity       string `yaml:"parity"`
	Stopbits     int    `yaml:"stopbits"`
}

func getSize(c SerialConf) byte {
	return byte(c.Databits)
}

func getParity(c SerialConf) serial.Parity {
	switch c.Parity {
	case "None":
		return serial.ParityNone
	default:
		log.Printf("default parity selected: None")
		return serial.ParityNone
	}
}

func getStopbits(c SerialConf) serial.StopBits {
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

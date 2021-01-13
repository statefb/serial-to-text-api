package myserial

import (
	"app/server/config"
	"log"
)

type SerialPort interface {
	Read([]byte) (int, error)
	Readline(n int) (string, error)
}

func GetSerialPort(mock bool) SerialPort {
	c := GetConf(config.SerialConfPath, true)
	if mock {
		log.Printf("using dummy serial port.")
		return NewDummySerialPort()
	} else {
		return NewCustomSerialPort(c)
	}
}

package myserial

import (
	"app/server/config"
	"log"
)

type SerialPort interface {
	Read([]byte) (int, error)
	Readline(n int) (string, error)
	Close() error
}

func GetSerialPort() SerialPort {
	// c := GetConf(config.SerialConfPath, true)
	s := config.NewConf()
	if s.Serial.Mock {
		log.Printf("using dummy serial port.")
		return NewDummySerialPort()
	} else {
		sc := s.GetSerialConf()
		return NewCustomSerialPort(sc, s.Serial.EndSignature)
	}
}

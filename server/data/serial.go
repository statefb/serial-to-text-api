package data

import (
	"context"

	"local.packages/config"
)

type SerialPort interface {
	Open() error
	Readline(context.Context)
	Close()
}

func GetSerialPort(conf *config.Conf) SerialPort {
	mock := conf.Serial.Mock
	if !mock {
		return NewCustomSerial(conf)
	} else {
		return NewMockSerial(conf)
	}
}

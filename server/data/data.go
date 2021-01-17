package data

import (
	"app/server/gen/models"
	"app/server/myserial"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

// serial port
var sp myserial.SerialPort

// declare memory to keep received data
var data []models.CollectedData

// channel to communicate serial goroutine
var dchan chan models.CollectedData

// channel size (i.e. max buffer size to keep)
var cSize int = 10

// buffer size to read serial
var bSize int = 128

var opened = false

func Start() {
	// start collecting data from serial
	go func() {
		for i := 0; i < cSize; i++ {
			val, err := sp.Readline(bSize)
			if err != nil {
				panic(err)
			}

			record := models.CollectedData{
				Timestamp: conv.DateTime(strfmt.DateTime(time.Now())),
				Value:     swag.String(val),
			}
			dchan <- record
		}
	}()
}

func Open() error {
	if !opened {
		Reset()
		// open channel and serial port
		dchan = make(chan models.CollectedData, 5)
		sp = myserial.GetSerialPort()
		return nil
	}
	return nil
}

func Close() error {
	Reset()
	// close channel and serial port
	close(dchan)
	err := sp.Close()
	if err != nil {
		return err
	}
	return nil
}

func Reset() error {
	data = []models.CollectedData{}
	return nil
}

func GetData() []models.CollectedData {
	timeout := time.After(time.Millisecond * time.Duration(1))
	for {
		select {
		case r := <-dchan:
			data = append(data, r)
		case <-timeout:
			return data
		}
	}
}

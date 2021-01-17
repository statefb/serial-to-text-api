package data

import (
	"app/server/gen/models"
	"app/server/myserial"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

// declare memory to keep received data
var data []*models.CollectedData

// channel to communicate serial goroutine
var dchan chan *models.CollectedData

// channel size (i.e. max buffer size to keep)
var cSize int = 10

// buffer size to read serial
var bSize int = 128

func Initialize() {
	// initialize
	data = nil
	dchan = make(chan *models.CollectedData, 5)
	s := myserial.GetSerialPort()

	go receive(s)
}

// func receive(s myserial.SerialPort) {
// 	val, err := s.Readline(bSize)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	record := &models.CollectedData{
// 		Timestamp: conv.DateTime(strfmt.DateTime(time.Now())),
// 		Value:     swag.String(val),
// 	}
// 	dchan <- record
// }

func Close() error {
	close(dchan)
	// TODO: close serial port
	return nil
}

func receive(s myserial.SerialPort) {
	for i := 0; i < cSize; i++ {
		val, err := s.Readline(bSize)
		if err != nil {
			panic(err)
		}

		record := &models.CollectedData{
			Timestamp: conv.DateTime(strfmt.DateTime(time.Now())),
			Value:     swag.String(val),
		}
		dchan <- record
	}
}

func GetData() []*models.CollectedData {
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

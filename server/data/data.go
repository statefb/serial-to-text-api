package data

import (
	"app/server/gen/models"
	"app/server/myserial"
	"log"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

// declare memory to have received data
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
	mock := true
	s := myserial.GetSerialPort(mock)

	for i := 0; i < cSize; i++ {
		go receive(s)
	}
}

func receive(s myserial.SerialPort) {
	str, err := s.Readline(bSize)
	if err != nil {
		log.Fatal(err)
	}
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}

	record := &models.CollectedData{
		Timestamp: conv.DateTime(strfmt.DateTime(time.Now())),
		Value:     swag.Float64(val),
	}
	dchan <- record
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

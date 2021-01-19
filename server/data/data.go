package data

import (
	"app/server/config"
	"app/server/gen/models"
	"app/server/myserial"
	"fmt"
	"log"
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
var cSize int

// buffer size to read serial
var bSize int

// whether port is opened
var opened = false

func Start() {
	s := config.NewConf().Serial
	cSize = s.MaxRecordSize
	bSize = s.BufferSize
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
			fmt.Printf("record: %+v\n", record)
			dchan <- record
		}
	}()
}

func Open() error {
	if !opened {
		data = []models.CollectedData{}
		// open channel and serial port
		dchan = make(chan models.CollectedData, 5)
		sp = myserial.GetSerialPort()
		opened = true
		return nil
	}
	return nil
}

func Close() error {
	if opened {
		log.Printf("return data to empty.")
		data = []models.CollectedData{}
		// close channel and serial port
		log.Printf("close channel.")
		close(dchan)
		log.Printf("close port.")
		err := sp.Close()
		if err != nil {
			return err
		}
		opened = false
	}

	return nil
}

func Reset() error {
	err := Close()
	if err != nil {
		return err
	}
	err = Open()
	if err != nil {
		return err
	}
	return nil
}

func GetData() []models.CollectedData {
	timeout := time.After(time.Millisecond * time.Duration(1))
	for {
		select {
		case r := <-dchan:
			fmt.Printf("r: %+v\n", r)
			data = append(data, r)
		case <-timeout:
			fmt.Printf("data: %+v\n", data)
			return data
		}
	}
}

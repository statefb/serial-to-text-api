package data

import (
	"app/server/config"
	"app/server/gen/models"
	"context"
	"fmt"
	"log"
)

var ctx context.Context
var cancel context.CancelFunc

var data []models.CollectedData

func Start() {
	ctx = context.Background()
	ctx, cancel = context.WithCancel(ctx)
	// fmt.Printf("number of goroutine: %d\r\n", runtime.NumGoroutine())
	go collect(ctx)
}

func collect(ctx context.Context) {
	conf := config.NewConf()
	s := GetSerialPort(conf)
	ctxSerial, cancelSerial := context.WithCancel(ctx)
	defer cancelSerial()
	err := s.Open()
	if err != nil {
		panic(err)
	}

	for i := 0; i < conf.Serial.MaxRecordSize; i++ {
		select {
		case <-ctx.Done():
			log.Printf("collection stopped.")
			return
		default:
			fmt.Printf("%d th reading\r\n", i)
			s.Readline(ctxSerial)
		}
	}

	fmt.Printf("Max record size reached. Stop data collection.\r\n")
	s.Close()
}

func Stop() error {
	cancel()
	return nil
}

func Reset() error {
	cancel()
	data = []models.CollectedData{}
	return nil
}

func GetData() []models.CollectedData {
	return data
}

package data

import (
	"app/server/config"
	"app/server/gen/models"
	"context"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

type MockSerial struct{}

func NewMockSerial(conf *config.Conf) *MockSerial {
	return &MockSerial{}
}

func (s *MockSerial) Open() error {
	return nil
}

func (s *MockSerial) Close() {
	return
}

func (s *MockSerial) Readline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("stop dummy serial port.")
			return
		default:
			time.Sleep(time.Millisecond * 1000)
			setDummyRecord()
			return
		}
	}
}

func setDummyRecord() {
	record := models.CollectedData{
		Timestamp: conv.DateTime(strfmt.DateTime(time.Now())),
		Value:     swag.String("this is dummy serial signal."),
	}
	data = append(data, record)
}

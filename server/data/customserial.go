package data

import (
	"context"
	"log"
	"strings"
	"time"

	"local.packages/config"
	"local.packages/gen/models"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
	"github.com/tarm/serial"
)

type CustomSerial struct {
	port   *serial.Port
	config *serial.Config
	end    string
	que    que
	buffer int
}

func NewCustomSerial(conf *config.Conf) *CustomSerial {
	sc := conf.GetSerialConf()
	return &CustomSerial{
		config: sc,
		end:    conf.Serial.EndSignature,
		buffer: conf.Serial.BufferSize,
	}
}

func (s *CustomSerial) Open() error {
	p, err := serial.OpenPort(s.config)
	s.port = p
	return err
}

func (s *CustomSerial) Close() {
	s.port.Close()
	log.Printf("closed port.")
}

func (s *CustomSerial) Readline(ctx context.Context) {
	signal := ""
	for {
		select {
		case <-ctx.Done():
			log.Printf("stop reading from port.")
			s.Close()
			return
		default:
			// try to fetch from que
			x, err := s.que.pop()
			if err == nil {
				log.Printf("pop from que")
				log.Printf(x.(string))
				setRecord(x.(string))
				return
			}

			val := read(s.port)
			signal = signal + val
			if strings.Contains(signal, s.end) {
				// split by end signature
				splitted := strings.Split(signal, s.end)
				for _, sp := range splitted {
					if sp == "" {
						// ignore empty string
						continue
					}
					s.que.push(sp) // hold on que
				}
				x, _ = s.que.pop() // return first element
				log.Printf("read from serial")
				log.Printf(x.(string))
				setRecord(x.(string))
				return
			}
		}
	}

}

func read(p *serial.Port) string {
	buf := make([]byte, 128)
	n, err := p.Read(buf)
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	return string(buf[:n])
}

func setRecord(s string) {
	record := models.CollectedData{
		Timestamp: conv.DateTime(strfmt.DateTime(time.Now())),
		Value:     swag.String(s),
	}
	// fmt.Printf("setting following record: %+v\n", record)
	data = append(data, record)
}

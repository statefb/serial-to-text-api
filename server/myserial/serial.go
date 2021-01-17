package myserial

import (
	"log"
	"strings"

	"github.com/tarm/serial"
)

var q que
var end string

// SerialPort implementation (wrapper of serial.Port)
type CustomSerialPort struct {
	port *serial.Port
}

func NewCustomSerialPort(c *serial.Config, e string) *CustomSerialPort {
	end = e
	p, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
	return &CustomSerialPort{port: p}
}

// SerialPort interface implementation
func (s *CustomSerialPort) Read(b []byte) (int, error) {
	n, err := s.port.Read(b)
	return n, err
}

// SerialPort interface implementation
func (s *CustomSerialPort) Readline(size int) (string, error) {
	// try to fetch from que
	x, err := q.pop()
	if err == nil {
		return x.(string), nil
	}
	// if que has no element, read from serial port.
	// loop until get end signature
	str := ""
	for {
		buf := make([]byte, size)
		n, err := s.Read(buf)
		if err != nil {
			return "", err
		}
		str = str + string(buf[:n])
		log.Printf(str)
		if strings.Contains(str, end) {
			break
		}
	}
	// split by end signature
	splitted := strings.Split(str, end)
	for _, sp := range splitted {
		log.Printf(sp)
		q.push(sp) // hold on que
	}
	x, _ = q.pop() // return first element
	return x.(string), nil
}

func (s *CustomSerialPort) Close() (err error) {
	return s.port.Close()
}

package myserial

import (
	"log"
	"strings"

	"github.com/tarm/serial"
)

// SerialPort implementation (wrapper of serial.Port)
type CustomSerialPort struct {
	port *serial.Port
}

func NewCustomSerialPort(c *serial.Config) *CustomSerialPort {
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
	buf := make([]byte, size)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	str := string(buf[:n])
	// split by return code
	splitted := strings.Split(str, "\r\n")
	log.Printf("splitted signal: \r\n")
	for i, sp := range splitted {
		log.Printf(string(i) + ": " + sp)
	}
	return splitted[0], err
}

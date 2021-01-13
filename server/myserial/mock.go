package myserial

// "time"

type DummySerialPort struct{}

func NewDummySerialPort() *DummySerialPort {
	return &DummySerialPort{}
}

// SerialPort interface implementation
func (s *DummySerialPort) Read(p []byte) (int, error) {
	// time.Sleep(time.Second)
	p = []byte("25.123")
	return 8, nil
}

// SerialPort interface implementation
func (s *DummySerialPort) Readline(n int) (string, error) {
	return "25.234", nil
}

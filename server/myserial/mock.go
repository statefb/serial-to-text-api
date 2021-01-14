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
	return "H234---Value:123431--\r\nH234---Weight:214211--", nil
}

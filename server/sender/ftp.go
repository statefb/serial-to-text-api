package sender

import (
	"app/server/converter"
	"app/server/data"
	"bytes"

	"github.com/jlaffaye/ftp"
)

type FTPSender struct {
	TargetPath string
	Uri        string
	Name       string
	Password   string
}

func NewFTPSender() *FTPSender {
	return &FTPSender{
		TargetPath: "test/test.txt",
		Uri:        "localhost:21",
		Name:       "anonymous",
		Password:   "password",
	}
}

func (s *FTPSender) Send() error {
	client, err := ftp.Connect(s.Uri)
	if err != nil {
		return err
	}
	defer client.Quit()

	err = client.Login(s.Name, s.Password)
	if err != nil {
		return err
	}

	converter := converter.NewJsonConverter(
		data.GetKeyData(), data.GetData(),
	)
	content, err := converter.Convert()
	if err != nil {
		return err
	}

	reader := bytes.NewReader(content)

	err = client.Stor(s.TargetPath, reader)
	if err != nil {
		return err
	}

	return nil
}

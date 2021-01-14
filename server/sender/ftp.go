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
		TargetPath: "test.json",
		Uri:        "ftp:21",
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

	kd := data.GetKeyData()
	cd := data.GetData()
	converter := converter.NewJsonConverter(
		kd, cd,
	)
	content, err := converter.Convert()
	if err != nil {
		return err
	}
	reader := bytes.NewReader(content)

	// For consistency, generate temporary file then rename.
	// this implementation expects that the ftp server does not handle file while transportation.
	// server should not touch files which filename contains "temp".
	tp := "temp" + randName(15) + ".json"
	err = client.Stor(tp, reader)
	if err != nil {
		return err
	}
	err = client.Rename(tp, s.TargetPath)
	if err != nil {
		return err
	}

	return nil
}

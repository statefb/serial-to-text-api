package sender

import (
	"bytes"

	"local.packages/config"
	"local.packages/converter"
	"local.packages/data"
	"local.packages/gen/models"

	"github.com/jlaffaye/ftp"
)

type FTPSender struct {
	TargetPath string
	Uri        string
	Name       string
	Password   string
}

func NewFTPSender() *FTPSender {
	s := config.NewConf()
	ftp := s.Sender.Ftp
	return &FTPSender{
		TargetPath: ftp.Path,
		Uri:        ftp.Uri,
		Name:       ftp.Name,
		Password:   ftp.Password,
	}
}

func (s *FTPSender) Send(body []*models.CollectedData) error {
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
	cd := convertCollectData(body)
	converter := converter.NewJsonConverter(
		&kd, &cd,
	)
	content, err := converter.Convert()
	if err != nil {
		return err
	}
	reader := bytes.NewReader(content)

	// For consistency, generate temporary file then rename.
	// this implementation expects that the ftp server does not handle file while transfer.
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

func (s *FTPSender) SendAll() error {
	d := data.GetData()
	dat := make([]*models.CollectedData, len(d))
	for i := 0; i < len(d); i++ {
		dat[i] = &d[i]
	}
	return s.Send(dat)
}

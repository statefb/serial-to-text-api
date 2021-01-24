package sender

import (
	"os"

	"local.packages/config"
	"local.packages/converter"
	"local.packages/data"
)

type LocalSender struct {
	TargetPath string
}

func NewLocalSender() *LocalSender {
	s := config.NewConf()
	local := s.Sender.Local
	return &LocalSender{
		TargetPath: local.Path,
	}
}

func (s *LocalSender) Send() error {
	kd := data.GetKeyData()
	cd := data.GetData()
	converter := converter.NewJsonConverter(
		&kd, &cd,
	)
	content, err := converter.Convert()
	if err != nil {
		return err
	}

	file, err := os.Create(s.TargetPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

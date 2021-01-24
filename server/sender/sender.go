package sender

import (
	"local.packages/config"

	"github.com/pkg/errors"
)

type Sender interface {
	Send() error
}

func GetSender() (Sender, error) {
	method := config.NewConf().Sender.Method
	switch method {
	case "ftp":
		return NewFTPSender(), nil
	case "local":
		return NewLocalSender(), nil
	}

	return nil, errors.Errorf("following method is not implemented: %s", method)
}

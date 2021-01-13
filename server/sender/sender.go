package sender

import "github.com/pkg/errors"

type Sender interface {
	Send() error
}

func GetSender() (Sender, error) {
	stype := "ftp"
	if stype == "ftp" {
		return NewFTPSender(), nil
	}
	return nil, errors.Errorf("following method is not implemented: %s", stype)
}

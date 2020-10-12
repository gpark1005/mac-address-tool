package params

import (
	"errors"
	"net"
)

type Options struct {
	ApiKey     string `flag:"k"`
	MacAddress string `flag:"a"`
}

func (o Options) Valid() error {
	_, err := net.ParseMAC(o.MacAddress)
	if err != nil {
		err := errors.New("Please enter a valid mac address")
		return err
	}

	return nil
}

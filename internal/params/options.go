package params

import (
	"errors"
	"net"
)

type Options struct {
	ApiKey     string `flag:"k"`
	MacAddress string `flag:"a"`
	Format     string `flag:"f"`
}

func (o Options) Valid() error {
	_, err := net.ParseMAC(o.MacAddress)
	if err != nil {
		err := errors.New("Please enter a valid mac address")
		return err
	}

	if o.Format != "" && o.Format != "json" && o.Format != "xml" && o.Format != "csv" {
		err := errors.New("Please enter a valid return format, for example: json")
		return err
	}

	return nil
}

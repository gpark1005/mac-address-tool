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

	// Add delimiters if needed
	addr := o.MacAddress
	if len(addr) == 12 {
		for i := 2; i < len(addr); i += 3 {
			addr = addr[:i] + ":" + addr[i:]
		}
	}

	// Validate MAC address
	_, err := net.ParseMAC(o.MacAddress)
	if err != nil {
		err := errors.New("Please enter a valid mac address")
		return err
	}

	// Validate format argument
	if o.Format != "" && o.Format != "json" && o.Format != "xml" && o.Format != "csv" {
		err := errors.New("Please enter a valid return format, for example: json")
		return err
	}

	if o.ApiKey == "" {
		err := errors.New("Please enter a valid API key")
		return err
	}

	return nil
}

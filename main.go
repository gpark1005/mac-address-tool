package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mac-address-tool/internal/params"
	"net/http"
	"os"
	"time"
)

func main() {

	var opt params.Options

	flag.StringVar(&opt.ApiKey, "k", "", "(Required) The API request key")
	flag.StringVar(&opt.MacAddress, "a", "", "(Required) The MAC address to query")
	flag.StringVar(&opt.Format, "f", "", "(Optional) The format of the return data. Acceptable arguments are: json, xml, or csv. If no argument is specified, "+
		"only the vendor name will be returned")
	flag.Parse()

	if err := opt.Valid(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", "https://api.macaddress.io/v1", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("apiKey", opt.ApiKey)
	q.Add("search", opt.MacAddress)
	q.Add("output", opt.Format)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
}

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

	flag.StringVar(&opt.ApiKey, "k", "", "The API request key")
	flag.StringVar(&opt.MacAddress, "a", "", "The MAC address to query")
	flag.Parse()

	req, err := http.NewRequest("GET", "https://api.macaddress.io/v1", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("apiKey", opt.ApiKey)
	q.Add("search", opt.MacAddress)
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

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	for {
		test_func()
	}
}

func test_func() {
	u, _ := url.ParseRequestURI("https://golang.org")
	urlStr := fmt.Sprintf("%v", u)
	req, _ := http.NewRequest("GET", urlStr, nil)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	cli := &http.Client{Transport: tr}
	response, err := cli.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

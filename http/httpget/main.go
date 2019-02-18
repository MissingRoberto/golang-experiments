package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	address := flag.String("address", "localhost:8080", "address of the http server")
	flag.Parse()

	fmt.Printf("Calling %s\n", *address)
	req, err := http.Get(fmt.Sprintf("http://%s", *address))
	if err != nil {
		log.Fatal("request failed", err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("reading body failed", err)
	}
	fmt.Println(string(body))
}

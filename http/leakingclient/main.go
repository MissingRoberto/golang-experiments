package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	address := flag.String("address", "localhost:8080", "address of the http server")
	numRequests := flag.Int("requests", 1, "number of requests performed")
	interval := flag.Int("interval", 500, "interval between requests in milliseconds")
	sleep := flag.Int("sleep", 30, "sleep se between requests")

	flag.Parse()

	var clients []http.Client
	for i := 0; i < *numRequests; i++ {
		fmt.Printf("Calling %s\n", *address)
		client := http.Client{}
		clients = append(clients, client)
		request, err := http.NewRequest("GET", fmt.Sprintf("http://%s/%d", *address, i), nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := client.Do(request)
		if err != nil {
			log.Fatal("request failed", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			log.Fatal("reading body failed", err)
		}
		fmt.Println(string(body))
		fmt.Printf("Waiting %v milliseconds\n", time.Duration(*interval)*time.Millisecond)
		time.Sleep(time.Duration(*interval) * time.Millisecond)
	}
	fmt.Printf("Sleeping %v seconds\n", time.Duration(*sleep)*time.Second)
	time.Sleep(time.Duration(*sleep) * time.Second)
}

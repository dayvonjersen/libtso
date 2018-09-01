package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	payload, err := json.MarshalIndent(struct {
		HelloWorld string `json:"hello_world"`
	}{
		HelloWorld: "Hello, world!",
	}, "", "  ")
	checkErr(err)
	log.Println(string(payload))

	req, err := http.NewRequest("POST", "http://localhost", bytes.NewBuffer(payload))
	checkErr(err)

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkErr(err)

	/*
		POST / HTTP/1.1
		Host: localhost
		User-Agent: Go-http-client/1.1
		Content-Length: 36
		Content-Type: application/json
		Accept-Encoding: gzip

		{
		  "hello_world": "Hello, world!"
		}
	*/

	log.Println("[resp] Status: ", resp.Status)
	log.Println("[resp] Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	log.Println("[resp] Body:   ", string(body))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

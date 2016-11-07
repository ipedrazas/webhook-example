package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AnExample struct {
	Name string
	Age  int
}

type Hook struct {
	Data          AnExample
	Url           string
	TlsSkipVerify bool
	Debug         bool
}

func main() {
	data := AnExample{
		Name: "my name",
		Age:  42,
	}

	hook := &Hook{
		Url:           "http://localhost:8080",
		TlsSkipVerify: false,
		Debug:         true,
		Data:          data,
	}
	sendHook(*hook)
}

func sendHook(hook Hook) {

	client := &http.Client{}
	if hook.TlsSkipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	data := new(bytes.Buffer)
	json.NewEncoder(data).Encode(hook.Data)
	req, err := http.NewRequest("POST", hook.Url, data)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if hook.Debug {
		log.Println("response Status:", resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("response Body:", string(body))
	}
}

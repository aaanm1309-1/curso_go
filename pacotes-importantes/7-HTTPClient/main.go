package main

import (
	"io"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Second * 10}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	print(string(body))

}

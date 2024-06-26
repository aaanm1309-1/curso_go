package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// req, err := http.Get("https://www.google.com")
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	// countTime := time.Now().Add(1 * time.Second)
	countTime := time.Now().Add(60 * time.Second)
	count := 0

	// for count < 1000 {
	for countTime.After(time.Now()) {

		number := randGenerator.Intn(99) + 1
		cep := "012" + fmt.Sprintf("%02d", number) + "000"

		req, err := http.Get("http://localhost:8084/?cep=" + cep)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		_, err = io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		// res, err := io.ReadAll(req.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// println(string(res))
		count++
	}
	println(count)

}

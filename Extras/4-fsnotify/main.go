package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	MarshalConfig("config.json")
	fmt.Println(config)

	done := make(chan bool)
	go func() {
		for {

			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					fmt.Println("modified file:", event.Name)
					fmt.Println(config)

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	// err2 := json.Unmarshal(data, &config)
	fmt.Println(data)
	err2 := json.NewDecoder(bytes.NewReader(data)).Decode(&config)
	if err2 != nil {
		panic(err2)
	}

}

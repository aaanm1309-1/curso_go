package main

import "br.com.adrianomenezes/cursogo/9-Api/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)

}

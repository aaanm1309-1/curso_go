package main

import (
	"fmt"
)

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("Panic1 recovery")
			}
			if r == "panic2" {
				fmt.Println("Panic2 recovery")
			}
		}

	}()

	panic1()
	panic2()
}

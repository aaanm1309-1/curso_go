package main

import "fmt"

func main() {

	done := make(chan bool)

	values := []string{"a", "b", "c", "D", "E", "F", "G", "H", "I", "J"}

	x := 10
	for i := range x {
		fmt.Println(i)
	}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}

}

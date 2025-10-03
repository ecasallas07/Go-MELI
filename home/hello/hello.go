package main

import (
	"fmt"

	"example.com/greetings"
)

// This return module local
func main(){
	message:= greetings.Hello("Juan")
	fmt.Println(message)
}

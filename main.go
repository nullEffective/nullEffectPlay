package main

import (
	"fmt"
	"log"
)

func channel() {

	messages := make(chan string, 2)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func main() {
	log.Println("foo")

}

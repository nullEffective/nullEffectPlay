package main

import (
	"fmt"
	"github.com/nulleffect/nulleffectplay/effectserver"
	"log"
)

//github.com/nulleffect/nulleffectplay/effetserver

type Bar struct {
	name string
}

type test struct {
	name string
}

func name(i []byte) int {
	fmt.Println(i)
	return 0
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func channel() {

	messages := make(chan string, 2)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func main() {
	log.Println("foo")
	fmt.Println("fooxechoxxx")

	go f("goroutine")
	effectserver.Init()

}

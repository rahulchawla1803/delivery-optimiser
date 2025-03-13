package main

import (
	// "github.com/rahulchawla1803/delivery-optimiser/internal/hello"
	"log"

	"github.com/rahulchawla1803/delivery-optimiser/internal/core"
)

func main() {
	// hello.SayHello()

	err := core.Run()
	if err != nil {
		log.Fatal(err)
	}
}

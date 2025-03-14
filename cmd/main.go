package main

import (
	"log"

	"github.com/rahulchawla1803/delivery-optimiser/internal/core"
)

func main() {
	err := core.Run()
	if err != nil {
		log.Fatal(err)
	}
}

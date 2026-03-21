package main

import (
	"log"

	"github.com/Meduzz/commando"
)

func main() {
	err := commando.Execute()

	if err != nil {
		log.Fatalf("There was an error: %v", err)
	}
}

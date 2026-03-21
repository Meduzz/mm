package main

import (
	"log"

	"github.com/Meduzz/commando"
	_ "github.com/Meduzz/mm/cmd"
)

func main() {
	err := commando.Execute()

	if err != nil {
		log.Fatalf("There was an error: %v", err)
	}
}

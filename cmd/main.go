package main

import (
	"city-temp/internal/script"
	"log"
)

func main() {
	err := script.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = script.RunScript()
	if err != nil {
		log.Fatal(err)
	}
}

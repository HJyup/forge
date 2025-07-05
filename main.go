package main

import (
	"log"

	"github.com/HJyup/forge/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

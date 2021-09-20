package main

import (
	"github.com/r3m25/api-rest-golang/bd/configuration"
	"github.com/r3m25/api-rest-golang/handlers"
	"log"
)

func main() {
	if !configuration.ConnectionIsSuccess() {
		log.Fatal("Connection dont exist")
		return
	}

	handlers.Handlers()
}

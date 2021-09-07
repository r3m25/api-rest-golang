package main

import (
	"github.com/r3m25/api-rest-golang/bd"
	"github.com/r3m25/api-rest-golang/handlers"
	"log"
)

func main() {
	if !bd.ConnectionIsSuccess() {
		log.Fatal("Connection dont exist")
		return
	}

	handlers.Handlers()
}

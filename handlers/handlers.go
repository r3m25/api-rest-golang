package handlers

import (
	"github.com/gorilla/mux"
	"github.com/r3m25/api-rest-golang/middlew"
	"github.com/r3m25/api-rest-golang/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

//Handlers to set port, listen server
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckBd(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}

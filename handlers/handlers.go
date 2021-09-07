package handlers

import(
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

//Handlers to set port, listen server
func Handlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}

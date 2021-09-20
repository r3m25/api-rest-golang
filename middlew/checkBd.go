package middlew

import (
	"github.com/r3m25/api-rest-golang/bd/configuration"
	"net/http"
)

//CheckBd method to valid connection with database
func CheckBd(next http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		if !configuration.ConnectionIsSuccess() {
			http.Error(writer, "Connection was lost", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}

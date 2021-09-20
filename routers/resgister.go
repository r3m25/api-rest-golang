package routers

import (
	"encoding/json"
	"github.com/r3m25/api-rest-golang/bd/repository"
	"github.com/r3m25/api-rest-golang/models"
	"net/http"
)

//Register function which going to create a new user in database
func Register(w http.ResponseWriter, request *http.Request) {

	var user models.User
	// body is a stream read only, can read only one time.
	err := json.NewDecoder(request.Body).Decode(&user)

	if err!= nil {
		http.Error(w, "Has occurred an error " + err.Error(), http.StatusInternalServerError)
		return
	}

	if len(user.Mail) == 0 {
		http.Error(w, "Mail can't empty", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	_, exist, _ := repository.CheckIfExistUser(user.Mail)

	if exist {
		http.Error(w, "User already exist", http.StatusBadRequest)
		return
	}

	_, status, err := repository.SaveUser(user)
	if err != nil {
		http.Error(w, "Has occurred an error when tried to save a new user " + err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "User wasn't saved " + err.Error(), http.StatusBadRequest)
		return
	}

}
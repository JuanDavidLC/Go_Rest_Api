package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanDavidLC/Go_Rest_Api/models"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello there %s", "visitor")
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(users)

}

func addUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return

	}

	err = models.CreateUser(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "User was added")
}

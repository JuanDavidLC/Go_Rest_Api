package server

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(users)

}

func addUser(w http.ResponseWriter, r *http.Request) {

	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return

	}

	users = append(users, user)
	fmt.Fprintf(w, "User was added")
}

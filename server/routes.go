package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {

	http.HandleFunc("/", index)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			getUsers(w, r)

		case http.MethodPost:
			addUser(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Methos not allowes")
			return
		}

	})

}

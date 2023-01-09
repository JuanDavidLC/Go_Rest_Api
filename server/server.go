package server

import "net/http"

type User struct {
	Name     string
	Lastname string
	Age      int
}

var users []*User

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{

		Addr: addr,
	}

}

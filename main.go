package main

import (
	"github.com/JuanDavidLC/Go_Rest_Api/server"
)

func main() {

	srv := server.New(":8080")

	err := srv.ListenAndServe()

	if err != nil {

		panic(err)

	}
}

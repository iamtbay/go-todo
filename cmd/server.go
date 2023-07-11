package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/iamtbay/todo/routes"
)

func StartServer() error {
	server := http.Server{
		Handler: routes.Routes(),
		Addr:    fmt.Sprintf(":%v", os.Getenv("PORT")),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("server couldn't started!")
	}
	return nil
}

package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	logrus.Infof("Listening on: %v", server.Addr)
	http.HandleFunc("/lint/", lintHandler)
	http.Handle("/", http.FileServer(http.Dir("static/")))
	server.ListenAndServe()
}

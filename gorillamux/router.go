package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// gRPC
	mx := http.NewServeMux()
	mx.HandleFunc("/outlists", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/outlists"))
	})
	mx.HandleFunc("/outlists/points", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/outlists/points"))
	})
	// MY
	router := mux.NewRouter()
	sub := router.PathPrefix("/gw").Subrouter()
	sub.HandleFunc("/outlists/points/att", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("gw : /outlists/points/att"))
	})
	sub.HandleFunc("/outlists/points", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("gw : /outlists/points"))
	})

	mu := http.NewServeMux()
	mu.Handle("/", mx)
	mu.Handle("/gw/", router)
	server := &http.Server{
		Addr: "0.0.0.0:5000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mu, // Pass our instance of gorilla/mux in.
	}

	fmt.Println("starting server")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

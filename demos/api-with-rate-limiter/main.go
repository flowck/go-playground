package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Use httprate middleware to limit requests give the following rules:
	// 10 requests per IP per minute;
	r.Use(httprate.LimitByIP(10, 1*time.Minute))

	fmt.Println("Server is running at port 3000")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		panic(err)
	}
}

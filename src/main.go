package main

import (
	"log"
	"net/http"

	"example.com/src/app"
)

func main() {
	a := app.Initialize("postgres://user:pass@localhost:5433/connect?sslmode=disable")
	defer a.Close()

	mux := http.NewServeMux()
	a.RegisterHandlers(mux)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/gen/go/proto/examplev1connect"
	"example.com/src/db"
	"example.com/src/infra"
	"example.com/src/service"
)

func main() {
	dbConn := infra.NewPostgresDB("postgres://user:pass@localhost:5433/connect?sslmode=disable")
	defer dbConn.Close()

	greetingRepo := db.NewBunGreetingRepository(dbConn)

	server := service.NewHelloServer(greetingRepo)

	mux := http.NewServeMux()
	path, handler := examplev1connect.NewHelloServiceHandler(server)
	mux.Handle(path, handler)

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

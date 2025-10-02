package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/gen/go/proto/examplev1connect"
	"example.com/src/infra"
	"example.com/src/middleware"
	"example.com/src/repository"
	"example.com/src/service"
	"github.com/bufbuild/connect-go"
)

func main() {
	dbConn := infra.NewPostgresDB("postgres://user:pass@localhost:5433/connect?sslmode=disable")
	defer dbConn.Close()

	// Repositories
	greetingRepo := repository.NewBunGreetingRepository(dbConn)
	userRepo := repository.NewBunUserRepository(dbConn)

	// Services
	helloServer := service.NewHelloServer(greetingRepo)
	userServer := service.NewUserServer(userRepo)

	mux := http.NewServeMux()

	// HelloService
	helloPath, helloHandler := examplev1connect.NewHelloServiceHandler(helloServer, connect.WithInterceptors(middleware.LoggingInterceptor()))
	mux.Handle(helloPath, helloHandler)

	// UserService
	userPath, userHandler := examplev1connect.NewUserServiceHandler(userServer, connect.WithInterceptors(middleware.LoggingInterceptor()))
	mux.Handle(userPath, userHandler)

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

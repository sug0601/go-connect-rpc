package app

import (
	"net/http"

	"example.com/gen/go/proto/protoconnect"
	"example.com/src/infra"
	"example.com/src/middleware"
	"example.com/src/repository"
	"example.com/src/service"
	"github.com/bufbuild/connect-go"
	"github.com/uptrace/bun"
)

type App struct {
	DB       *bun.DB
	Repos    *repository.Repositories
	Services *service.Services
}

func Initialize(dsn string) *App {
	db := infra.NewPostgresDB(dsn)
	repos := repository.NewRepositories(db)
	services := service.NewServices(repos)

	return &App{
		DB:       db,
		Repos:    repos,
		Services: services,
	}
}

func (a *App) Close() {
	a.DB.Close()
}

func (a *App) RegisterHandlers(mux *http.ServeMux) {
	interceptors := connect.WithInterceptors(
		middleware.ValidationInterceptor(),
		middleware.LoggingInterceptor(),
	)

	// HelloService
	helloPath, helloHandler := protoconnect.NewHelloServiceHandler(
		a.Services.HelloServer,
		interceptors,
	)
	mux.Handle(helloPath, helloHandler)

	// UserService
	userPath, userHandler := protoconnect.NewUserServiceHandler(
		a.Services.UserServer,
		interceptors,
	)
	mux.Handle(userPath, userHandler)
}

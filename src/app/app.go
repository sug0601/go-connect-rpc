package app

import (
	"net/http"

	"example.com/gen/go/proto/examplev1connect"
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
	helloPath, helloHandler := examplev1connect.NewHelloServiceHandler(a.Services.HelloServer, connect.WithInterceptors(middleware.LoggingInterceptor()))
	mux.Handle(helloPath, helloHandler)

	userPath, userHandler := examplev1connect.NewUserServiceHandler(a.Services.UserServer, connect.WithInterceptors(middleware.LoggingInterceptor()))
	mux.Handle(userPath, userHandler)
}

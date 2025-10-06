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

type Service struct {
    Server     interface{}
    NewHandler func(s interface{}, interceptors ...connect.HandlerOption) (string, http.Handler)
}

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

	services := []Service{
		{
			Server: a.Services.HelloServer,
			NewHandler: func(s interface{}, interceptors ...connect.HandlerOption) (string, http.Handler) {
				return protoconnect.NewHelloServiceHandler(s.(*service.HelloServer), interceptors...)
			},
		},
		{
			Server: a.Services.UserServer,
			NewHandler: func(s interface{}, interceptors ...connect.HandlerOption) (string, http.Handler) {
				return protoconnect.NewUserServiceHandler(s.(*service.UserServer), interceptors...)
			},
		},
		{
			Server: a.Services.ArticleServer,
			NewHandler: func(s interface{}, interceptors ...connect.HandlerOption) (string, http.Handler) {
				return protoconnect.NewArticleServiceHandler(s.(*service.ArticleServer), interceptors...)
			},
		},
	}

	for _, svc := range services {
		path, handler := svc.NewHandler(svc.Server, interceptors)
		mux.Handle(path, handler)
	}
}
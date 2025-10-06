package service

import "example.com/src/repository"

type Services struct {
	HelloServer *HelloServer
	UserServer  *UserServer
	ArticleServer  *ArticleServer
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		HelloServer: NewHelloServer(repos.Greeting),
		UserServer:  NewUserServer(repos.User),
		ArticleServer: NewArticleServer(repos.Article),
	}
}

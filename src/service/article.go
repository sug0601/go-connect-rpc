package service

import (
	"context"
	"fmt"

	examplev1 "example.com/gen/go/proto"
	"example.com/src/repository"
	"github.com/bufbuild/connect-go"
)

type ArticleServer struct {
	articleRepo repository.ArticleRepository
}

func NewArticleServer(repo repository.ArticleRepository) *ArticleServer {
	return &ArticleServer{articleRepo: repo}
}

func (s *ArticleServer) ListArticles(
	ctx context.Context,
	req *connect.Request[examplev1.ListArticlesRequest],
) (*connect.Response[examplev1.ListArticlesResponse], error) {

	articles, err := s.articleRepo.FindAll(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var protoArticles []*examplev1.ArticleSummary
	for _, a := range articles {
		protoArticles = append(protoArticles, &examplev1.ArticleSummary{
			Id:        a.ID,
			Title:     a.Title,
			Thumbnail: a.Thumbnail,
		})
	}

	res := connect.NewResponse(&examplev1.ListArticlesResponse{
		Articles: protoArticles,
	})
	return res, nil
}

func (s *ArticleServer) GetArticle(
	ctx context.Context,
	req *connect.Request[examplev1.GetArticleRequest],
) (*connect.Response[examplev1.GetArticleResponse], error) {

	article, err := s.articleRepo.FindByID(ctx, req.Msg.ArticleId)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("article not found"))
	}

	res := connect.NewResponse(&examplev1.GetArticleResponse{
		Article: &examplev1.Article{
			Id:        article.ID,
			Title:     article.Title,
			Thumbnail: article.Thumbnail,
			Content:   article.Content,
		},
	})
	return res, nil
}

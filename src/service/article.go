package service

import (
	"context"
	"fmt"
	"log"

	examplev1 "example.com/gen/go/proto"
	"example.com/src/helper"
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
	p := req.Msg.Pagination
	if p == nil {
		p = &examplev1.PaginationRequest{
			Page:     1,
			PageSize: 1,
		}
	}

	pagination := helper.NewPagination(int(p.Page), int(p.PageSize))

	articles, totalCount, err := s.articleRepo.FindAll(ctx, pagination.Offset, pagination.PageSize)
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

	totalPages := int32((totalCount + int64(p.PageSize) - 1) / int64(p.PageSize))

	res := connect.NewResponse(&examplev1.ListArticlesResponse{
		Articles: protoArticles,
		Pagination: &examplev1.PaginationResponse{
			Page:       p.Page,
			PageSize:   p.PageSize,
			TotalCount: int32(totalCount),
			TotalPages: totalPages,
		},
	})

	log.Printf("Returning response with %d articles", len(protoArticles))
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

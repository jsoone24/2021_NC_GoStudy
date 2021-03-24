package usecase

import "gin_V0.2/domain/model"

type ManageArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
	CreateNewArticle(title, content string) (*model.Article, error)
	DeleteArticleByID(id int) error
}

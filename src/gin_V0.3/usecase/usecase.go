package usecase

import "gin_V0.3/domain/model"

type ManageArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
	CreateNewArticle(title, content string, writerID int) (*model.Article, error)
	DeleteArticleByID(id int) error
}

type RegistrationUsecase interface {
	RegisterUser(name, pass string) (*model.User, error)
	MatchUser(name, pass string) (*model.User, error)
}

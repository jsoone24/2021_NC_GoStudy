package repository

import "gin_V0.3/domain/model"

type ArticleRepo interface {
	GetAll() ([]model.Article, error)
	GetByID(id int) (*model.Article, error)
	Create(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}

type UserRepo interface {
	GetByID(id int) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
}

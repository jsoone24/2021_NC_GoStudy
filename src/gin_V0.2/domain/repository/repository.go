package repository

import "gin_V0.2/domain/model"

type ArticleRepo interface {
	GetAll() ([]model.Article, error)
	GetByID(id int) (*model.Article, error)
	Create(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}

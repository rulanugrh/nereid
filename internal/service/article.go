package service

import (
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"github.com/rulanugrh/isonoe/internal/middleware"
	"github.com/rulanugrh/isonoe/internal/repository"
)

type ArticleInterface interface {
	Create(req domain.Article) (*web.ArticleCreate, error)
	GetById(id string) (*web.GetArticle, error)
	GetAll() (*[]web.GetArticle, error)
	Delete(id string) error
}

type article struct {
	repo repository.ArticleInterface
	validate middleware.ValidationInterface
}

func NewArticleService(repo repository.ArticleInterface) ArticleInterface {
	return &article{
		repo: repo,
		validate: middleware.NewValidation(),
	}
}

func (a *article) Create(req domain.Article) (*web.ArticleCreate, error) {
	err := a.validate.ValidateData(req)
	if err != nil {
		return nil, a.validate.ValidationMessage(err)
	}
	
	data, err := a.repo.Create(req)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	response := web.ArticleCreate{
		Title:   data.Title,
		Content: data.Content,
		Tags:    data.Tags,
		Author:  data.Author,
	}

	return &response, nil
}

func (a *article) GetById(id string) (*web.GetArticle, error) {
	data, err := a.repo.GetById(id)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	response := web.GetArticle{
		ID:          data.ID.String(),
		Title:       data.Title,
		Tags:        data.Tags,
		Author:      data.Author,
		Banner:      data.Banner,
		Content:     data.Content,
		CreatedAt:   data.CreatedAt,
		Description: data.Description,
		Conclusion:  data.Conclusion,
	}

	return &response, nil
}

func (a *article) GetAll() (*[]web.GetArticle, error) {
	data, err := a.repo.GetAll()
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	var response []web.GetArticle
	for _, dt := range *data {
		response = append(response, web.GetArticle{
			ID:          dt.ID.Hex(),
			Title:       dt.Title,
			Tags:        dt.Tags,
			Author:      dt.Author,
			Banner:      dt.Banner,
			Content:     dt.Content,
			CreatedAt:   dt.CreatedAt,
			Description: dt.Description,
		})
	}

	return &response, nil
}

func (a *article) Delete(id string) error {
	if err := a.repo.Delete(id); err != nil {
		return web.BadRequest(err.Error())
	}

	return nil
}

package service

import (
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"github.com/rulanugrh/isonoe/internal/repository"
)

type CommentInterface interface {
	// create comment
	CreateComment(req domain.CommentRequest) (*web.CommentCreate, error)
	// find all comment
	FindAllComment(articleID string) (*[]web.GetComment, error)
	// delete comment
	DeleteComment(id string, author string) error
}

type comment struct {
	repository repository.CommentInterface
}

func NewCommentService(repository repository.CommentInterface) CommentInterface {
	return &comment{repository: repository}
}

// create comment
func(c *comment) CreateComment(req domain.CommentRequest) (*web.CommentCreate, error) {
	data, err := c.repository.CreateComment(req)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	response := web.CommentCreate{
		Author: data.Author,
		Content: data.Content,
	}

	return &response, nil
}

// find all comment
func(c *comment) FindAllComment(articleID string) (*[]web.GetComment, error) {
	data, err := c.repository.GetCommentByArticle(articleID)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	var response []web.GetComment
	for _, value := range *data {
		comment := web.GetComment{
			ID: value.ID.Hex(),
			Author: value.Author,
			Content: value.Content,
		}

		response = append(response, comment)
	}

	return &response, nil
}

// delete comment
func(c *comment) DeleteComment(id string, author string) error {
	err := c.repository.DeleteComment(id, author)
	if err != nil {
		return web.BadRequest(err.Error())
	}

	return nil
}
package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"github.com/rulanugrh/isonoe/internal/service"
)

type ArticleInterface interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type article struct {
	service service.ArticleInterface
}

func NewArticleHandler(service service.ArticleInterface) ArticleInterface {
	return &article{
		service: service,
	}
}

func (a *article) Create(c *fiber.Ctx) error {
	var request domain.Article
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(500).JSON(web.InternalServerError("ERROR - Parsing body request"))
	}

	data, err := a.service.Create(request)
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return c.Status(201).JSON(web.Created("Success Create Article", data))
}

func (a *article) GetAll(c *fiber.Ctx) error {
	// parsing request for get all
	data, err := a.service.GetAll()
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	// return response
	return c.Render("index", fiber.Map{
		"Title": "Kyora | Blog",
		"Data":  data,
	})
}

func (a *article) Delete(c *fiber.Ctx) error {
	// parsing id parameter
	id := c.Params("id")

	// parsing id to service
	err := a.service.Delete(id)
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return c.Status(200).JSON(web.Success("success delete account", id))
}

func (a *article) GetById(c *fiber.Ctx) error {
	// parsing id parameter
	id := c.Params("id")

	data, err := a.service.GetById(id)
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return c.Render("article/index", fiber.Map{
		"Title":       data.Title,
		"Tags":        data.Tags,
		"Banner":      data.Banner,
		"Author":      data.Author,
		"Content":     data.Content,
		"CreatedAt":   data.CreatedAt,
		"ID":          data.ID,
		"Description": data.Description,
		"Conclusion":  data.Conclusion,
	})

}

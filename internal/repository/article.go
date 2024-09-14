package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/rulanugrh/isonoe/config"
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleInterface interface {
	// Create new article
	Create(req domain.Article) (*domain.Article, error)
	// Get article by id
	GetById(id string) (*domain.Article, error)
	// Find all article
	GetAll() (*[]domain.Article, error)
	// Delete article by id
	Delete(id string) error
}

type article struct {
	// collection
	coll *mongo.Collection
}

func NewArticleRepository(client *config.Connection, conf *config.App) ArticleInterface {
	return &article{
		coll: client.DB.Database(conf.Database.Name).Collection("articles"),
	}
}

// Create new article
func (a *article) Create(req domain.Article) (*domain.Article, error) {
	// Create new variable for decode response
	var response domain.Article
	// current time
	year, month, day := time.Now().Date()
	tm := fmt.Sprintf("%d %s %d", day, month, year)

	// create context with timeout with 10 second
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	// parsing into new request variabel
	request := domain.Article{
		Title:       req.Title,
		Content:     req.Content,
		Banner:      req.Banner,
		Author:      req.Author,
		Tags:        req.Tags,
		CreatedAt:   tm,
		UpdatedAt:   tm,
		Description: req.Description,
		Conclusion:  req.Conclusion,
	}

	// Insert new article
	data, err := a.coll.InsertOne(ctx, &request)
	if err != nil {
		return nil, web.ErrorLog("cannot insert into database")
	}

	err = a.coll.FindOne(ctx, bson.M{"_id": data.InsertedID}).Decode(&response)
	if err != nil {
		return nil, web.ErrorLog("cannot parsing response")
	}

	return &response, nil
}

// Get article by id
func (a *article) GetById(id string) (*domain.Article, error) {
	// parsing into new response
	var response domain.Article
	// create new context with timeout 10 seconds
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, web.ErrorLog("invalid id")
	}

	// get data in mongodb
	err = a.coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&response)
	if err != nil {
		return nil, web.ErrorLog("sorry data with this id not found")
	}

	return &response, nil
}

// Find all article
func (a *article) GetAll() (*[]domain.Article, error) {
	// create default variabel for parsing response
	var response []domain.Article
	// create new context with timeout 10 seconds
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	rows, err := a.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, web.ErrorLog("cannot find data in database")
	}

	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var article domain.Article
		if errs := rows.Decode(&article); errs != nil {
			return nil, web.ErrorLog("cannot parsing response")
		}

		response = append(response, article)
	}

	return &response, nil
}

// Delete article by id
func (a *article) Delete(id string) error {
	// create new context with timeout 10 seconds
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	// parsing id frm string
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return web.ErrorLog("invalid id")
	}

	_, err = a.coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return web.ErrorLog("cannot delete data in database")
	}

	return nil
}

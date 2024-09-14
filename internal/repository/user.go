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

type UserInterface interface {
	Create(req domain.UserRegister) (*domain.User, error)
	Login(req domain.UserLogin) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}

type user struct {
	client *mongo.Collection
}

func NewUserRepository(conn *config.Connection, conf *config.App) UserInterface {
	return &user{
		client: conn.DB.Database(conf.Database.Name).Collection("user"),
	}
}

func (u *user) Create(req domain.UserRegister) (*domain.User, error) {
	// response
	var response domain.User
	// current time
	t := time.Now()
	// parsing into user domain
	request := domain.User{
		Email:     req.Email,
		Password:  req.Password,
		Name:      req.Name,
		CreatedAt: primitive.NewDateTimeFromTime(t),
		UpdatedAt: primitive.NewDateTimeFromTime(t),
	}

	ctx, timeout := context.WithTimeout(context.Background(), 20*time.Second)
	defer timeout()

	data, err := u.client.InsertOne(ctx, &request)
	if err != nil {
		return nil, web.ErrorLog("failed to insert document user")
	}

	err = u.client.FindOne(ctx, bson.M{"_id": data.InsertedID}).Decode(&response)
	if err != nil {
		return nil, web.WarnLog("cannot find data with this id")
	}
	fmt.Println(response)
	return &response, nil
}

func (u *user) Login(req domain.UserLogin) (*domain.User, error) {
	// response
	var response domain.User

	// create context for 20 seconds
	ctx, timeout := context.WithTimeout(context.Background(), 20*time.Second)
	defer timeout()

	err := u.client.FindOne(ctx, bson.M{"email": req.Email}).Decode(&response)
	if err != nil {
		return nil, web.ErrorLog("sorry data not found")
	}

	return &response, nil
}

func (u *user) GetByEmail(email string) (*domain.User, error) {
	// response
	var response domain.User

	// create context for 20 seconds
	ctx, timeout := context.WithTimeout(context.Background(), 20*time.Second)
	defer timeout()

	err := u.client.FindOne(ctx, bson.M{"email": email}).Decode(&response)
	if err != nil {
		return nil, web.ErrorLog("sorry data not found")
	}

	return &response, nil
}

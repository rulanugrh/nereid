package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountCreate struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type AccountLogin struct {
	ID    primitive.ObjectID `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}

type GetAccount struct {
	ID    primitive.ObjectID `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}
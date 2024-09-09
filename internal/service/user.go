package service

import (
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"github.com/rulanugrh/isonoe/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	Register(req domain.UserRegister) (*web.AccountCreate, error)
	Login(req domain.UserLogin) (*web.AccountLogin, error)
	GetMe(email string) (*web.GetAccount, error)
}

type user struct {
	repository repository.UserInterface
}

func NewUserService(repo repository.UserInterface) UserInterface {
	return &user{repository: repo}
}

func(u *user) Register(req domain.UserRegister) (*web.AccountCreate, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, web.ErrorLog("something error while hashing password")
	}

	request := domain.UserRegister {
		Name: req.Name,
		Email: req.Email,
		Password: string(hashPassword),
	}

	data, err := u.repository.Create(request)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	return &web.AccountCreate{
		Name: data.Name,
		Email: data.Email,
	}, nil
}

func(u *user) Login(req domain.UserLogin) (*web.AccountLogin, error) {
	data, err := u.repository.Login(req)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	verify := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if verify != nil {
		return nil, web.WarnLog("Sorry you password is not matched")
	}

	return &web.AccountLogin{
		Name: data.Name,
		ID: data.ID,
		Email: data.Email,
	}, nil
}

func(u *user) GetMe(email string) (*web.GetAccount, error) {
	data, err := u.repository.GetByEmail(email)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	return &web.GetAccount{
		Name: data.Name,
		Email: data.Email,
		ID: data.ID,
	}, nil
}
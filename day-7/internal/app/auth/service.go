package auth

import (
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/middlewares"
	"alterra-agmc/day-6/internal/model"
	"alterra-agmc/day-6/internal/repository"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(model.User) (string, model.User, error)
	Register(model.User) (model.User, error)
}

type service struct {
	Repository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository}
}

func (s *service) Login(user model.User) (string, model.User, error) {
	data, err := s.Repository.FindByEmail(user.Email)
	if err != nil || reflect.ValueOf(data).IsZero() {
		return "", data, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password)); err != nil {
		return "", data, err
	}

	token, err := middlewares.CreateToken(int(data.ID))
	if err != nil {
		return "", data, err
	}

	return token, data, nil
}

func (s *service) Register(user model.User) (model.User, error) {
	user, err := s.Repository.Create(user)
	return user, err
}

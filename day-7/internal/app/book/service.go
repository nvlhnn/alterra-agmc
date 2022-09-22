package book

import (
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/model"
	"alterra-agmc/day-6/internal/repository"
)
type service struct {
	BookRepository repository.Book
}

type Service interface {
	FindAll() ([]model.Book, error)
	// FindByID(id uint) (model.User, error)
	// Update(id uint, user model.User) (model.User, error)
	// Delete(id uint) (model.User,error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) FindAll() ([]model.Book, error) {

	books, err := s.BookRepository.FindAll()
	return books, err

}

// func (s *service) FindByID(id uint) (model.User, error) {

// 	user, err := s.UserRepository.FindByID(id)
// 	return user, err
// }

// func (s *service) Update(id uint, updatedUser model.User) (model.User, error){

// 	user, err := s.FindByID(id)
// 	if err != nil {
// 		return user, err
// 	}

// 	if updatedUser.Email != "" {
// 		user.Email = updatedUser.Email
// 	}

// 	if updatedUser.Password != "" {
// 		user.Password = updatedUser.Password
// 	}

// 	user, err = s.UserRepository.Update(user)
// 	return user, err
// }

// func (s *service) Delete(id uint) (model.User, error){

// 	user, err := s.FindByID(id)
// 	if err != nil {
// 		return user, err
// 	}
// 	err = s.UserRepository.Delete(id)
// 	return user, err
// }
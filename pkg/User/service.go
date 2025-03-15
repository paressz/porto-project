package User

import "porto-project/pkg/model"

type Service interface {
	RegisterUser(user *model.User) error
	IsValidUsername(username string) bool
	IsValidEmail (email string) bool
	GetUser(email, password string) (*model.User, error)
//	UpdateUser(user *model.User) (*model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) RegisterUser(user *model.User) error {
	return s.repository.RegisterUser(user)
}

func (s *service) IsValidUsername(username string) bool {
	return s.repository.IsValidUsername(username)
}

func (s *service) IsValidEmail(email string) bool {
	return s.repository.IsValidEmail(email)
}

func (s *service) GetUser(email, password string) (*model.User, error) {
	return s.repository.GetUser(email, password)
}

//func (s *service) UpdateUser(user *model.User) (*model.User, error) {
//	return s.repository.UpdateUser(user)
//}

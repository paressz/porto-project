package User

import (
	"errors"
    "gorm.io/gorm"
	"porto-project/pkg/config/database"
	"porto-project/pkg/model"
    "porto-project/pkg/util/auth"
)

type Repository interface {
	RegisterUser(user *model.User) error
	IsValidUsername(username string) bool
	IsValidEmail (email string) bool
	GetUser(email, password string) (*model.User, error)
//	UpdateUser(user *model.User) (*model.User, error)
}

type repository struct {
	Db *gorm.DB
}

func NewRepository() Repository {
	db := database.Connect()
	return &repository{
		db,
	}
}

func (r repository) RegisterUser(user *model.User) error {
	if r.IsValidEmail(user.Email) || r.IsValidUsername(user.Username) {
		return errors.New("invalid Email or Username: Already Exist")
	}
	err := r.Db.
		Create(user).
		Error
	return err
}

func (r repository) IsValidUsername(username string) bool {
	var count int64
	r.Db.
		Model(&model.User{}).
		Where("username = ?", username).
		Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func (r repository) IsValidEmail(email string) bool {
	var count int64
	r.Db.
		Model(&model.User{}).
		Where("email = ?", email).
		Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func (r repository) GetUser(email, password string) (*model.User, error) {
	var user model.User
	err := r.Db.
		Where("email = ?", email).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	match := auth.ComparePasswordHash(password, user.Password)
    if !match {
		return nil, errors.New("Wrong password")
    }
	return &user, nil
}

//func (r repository) UpdateUser(user *model.User) (*model.User, error) {
//	panic("implement me")
//}
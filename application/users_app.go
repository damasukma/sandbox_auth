package application

import (
	"github.com/damasukma/sandbox_auth/domain/entity"
	"github.com/damasukma/sandbox_auth/domain/repository"
)

type (
	userApp struct {
		us repository.UserRepository
	}

	UserAppInterface interface {
		SaveUser(*entity.User) error
		EmailExist(email string) bool
		Auth(email, password string) (bool, error)
		Find() (*[]entity.User, error)
		FindID(id int) (*entity.User, error)
		Update(entity.User) error
		Delete(id int) error
	}
)

func (ua *userApp) SaveUser(user *entity.User) error {
	return ua.us.SaveUser(user)
}

func (ua *userApp) EmailExist(email string) bool {
	return ua.us.EmailExist(email)
}

func (ua *userApp) Auth(email, password string) (bool, error) {
	return ua.us.Auth(email, password)
}

func (ua *userApp) Find() (*[]entity.User, error) {
	return ua.Find()
}

func (ua *userApp) FindID(id int) (*entity.User, error) {
	return ua.FindID(id)
}

func (ua *userApp) Update(user entity.User) error {
	return ua.Update(user)
}

func (ua *userApp) Delete(id int) error {
	return ua.Delete(id)
}

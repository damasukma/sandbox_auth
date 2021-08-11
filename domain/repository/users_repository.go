package repository

import (
	"github.com/damasukma/sandbox_auth/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) error
	EmailExist(email string) bool
	Auth(email, password string) (bool, error)
	Find() (*[]entity.User, error)
	FindID(id int) (*entity.User, error)
	Update(user entity.User) error
	Delete(id int) error
}

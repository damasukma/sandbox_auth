package persistence

import (
	"github.com/damasukma/sandbox_auth/domain/entity"
	"github.com/damasukma/sandbox_auth/domain/repository"
	"gorm.io/gorm"
)

type (
	Repositories struct {
		User repository.UserRepository
		db   *gorm.DB
	}
)

func NewRepositories(connection *gorm.DB) *Repositories {
	return &Repositories{
		User: NewUsersRepository(connection),
		db:   connection,
	}
}

func (r *Repositories) Migrate() error {
	return r.db.AutoMigrate(&entity.User{})
}

package persistence

import (
	"errors"

	"github.com/damasukma/sandbox_auth/domain/entity"
	"github.com/damasukma/sandbox_auth/domain/repository"
	"github.com/damasukma/sandbox_auth/infrastructure/security"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

var _ repository.UserRepository = &UserRepo{}

func NewUsersRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) SaveUser(us *entity.User) error {
	if err := u.db.Debug().Create(&us).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) EmailExist(email string) bool {
	result := u.db.Debug().Where("email = ? ", email).First(&entity.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (u *UserRepo) Auth(email, password string) (bool, error) {
	var users *entity.User
	if err := u.db.Debug().Where("email = ? ", email).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}

	if err := security.Verify(users.Password, password); err != nil {
		return false, nil
	}
	return true, nil
}

func (u *UserRepo) Find() (*[]entity.User, error) {
	var users []entity.User
	if err := u.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserRepo) FindID(id int) (*entity.User, error) {
	var user entity.User
	if err := u.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) Update(user entity.User) error {

	if err := u.db.Debug().Select("email", "address").Where("id = ?", user.Id).Updates(&entity.User{Email: user.Email, Address: user.Address}).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) Delete(id int) error {
	var user entity.User
	if err := u.db.Debug().Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

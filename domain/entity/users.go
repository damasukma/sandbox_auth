package entity

import (
	"time"

	"github.com/damasukma/sandbox_auth/infrastructure/security"
	"gorm.io/gorm"
)

type (
	User struct {
		Id        int32      `json:"id"`
		Email     string     `json:"email"`
		Address   string     `json:"address"`
		Password  string     `json:"password"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}
)

func (u *User) Prepare() {
	u.CreatedAt = time.Now()
	*u.UpdatedAt = time.Now()
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hash, err := security.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

// func (u User) IsValid() (bool, *map[string]string) {
// 	var errMessages map[string]string

// 	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// 	if u.Email == "" || u.Email == "null" {
// 		errMessages[u.Email] = "email required"
// 	}

// 	if u.Address == "" || u.Address == "null" {
// 		errMessages[u.Address] = "address required"

// 	}

// 	if !re.MatchString(u.Email) {
// 		errMessages[u.Email] = "not email"

// 	}

// 	if len(u.Address) > 255 {
// 		errMessages[u.Address] = "address is too long"
// 	}

// 	if len(errMessages) > 0 {
// 		return false, &errMessages
// 	}

// 	return true, nil

// }

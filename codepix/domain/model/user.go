package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base     			`valid:"required"`
	Name     string 	`json:"name" valid:"notnull"`
	Email    string     `json:"email" valid:"notnull"`
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Name: name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

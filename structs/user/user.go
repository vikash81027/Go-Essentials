package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName string
	Birthdate string
	CreatedAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin {
		email : email,
		password: password,
		User: User {
			FirstName: "ADMIN",
			LastName: "ADMIN",
			Birthdate : "---",
	    CreatedAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	// can be used for validation

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("no empty fields")
	}
	
	return &User{
		FirstName: firstName,
		LastName: lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}

func (user *User) OutputUserDetail() {
	// go allows you to use struct pointer types without dereferencing it
	fmt.Println(user.FirstName, user.LastName, user.Birthdate, user.CreatedAt)
}

func (user *User) ClearuserName() {
	user.FirstName = ""
	user.LastName = ""
}

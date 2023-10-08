package model

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	CreationDate string `json:"creationDate"`
}

var (
	ErrNameInvalid     = errors.New("name is invalid")
	ErrIDInvalid       = errors.New("id is invalid")
	ErrEmailInvalid    = errors.New("email is invalid")
	ErrPasswordInvalid = errors.New("password is invalid")
)

type AddUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a AddUser) Validation() error {
	if a.Name == "" {
		return ErrNameInvalid
	}
	return nil
}

func UserOne(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, ErrNotFound
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u UpdateUser) Validation() error {
	switch {
	case u.ID == 0:
		return ErrIDInvalid
	case u.Name == "":
		return ErrNameInvalid
	case u.Password == "":
		return ErrPasswordInvalid
	case u.Email == "":
		return ErrEmailInvalid
	default:
		return nil
	}
}

func UsersAll(q string) ([]User, error) {
	if q != "" {
		var usersFiltered []User
		for _, user := range users {
			if user.Name == q {
				usersFiltered = append(usersFiltered, user)
			}
		}
		return usersFiltered, nil
	}
	return users, nil
}

func (u User) InsertUser() (int, error) {
	userMaxID++
	u.ID = userMaxID
	u.CreationDate = time.Now().Format("2006-01-02")
	users = append(users, u)
	return u.ID, nil
}

func DeleteUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (u User) UpdateUser() error {
	for i, user := range users {
		if user.ID == u.ID {
			users[i].Name = u.Name
			users[i].Password = u.Password
			users[i].Email = u.Email
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

var userMaxID = 3

var users = []User{
	{ID: 1, Name: "user1", Password: "pass1", Email: "user1@goggle.com", CreationDate: "2019-01-01"},
	{ID: 2, Name: "user2", Password: "pass2", Email: "user2@goggle.com", CreationDate: "2019-01-02"},
	{ID: 3, Name: "user3", Password: "pass3", Email: "user3@goggle.com", CreationDate: "2019-01-03"},
}

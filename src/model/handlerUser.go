package model

import (
	httperr "example/src/httpErr"
	"fmt"
)

type HandlerUser struct {
	users []User
}

func NewHandlerUser(users []User) *HandlerUser {
	return &HandlerUser{
		users: users,
	}
}

func (hu *HandlerUser) GetAll() ([]User, error) { return hu.users, nil }

func (hu *HandlerUser) CreateUser(user *User) (*User, error) {
	user.Id = UserId(len(hu.users) + 1)
	hu.users = append(hu.users, *user)
	return user, nil
}

func (hu *HandlerUser) GetById(userId UserId) (*User, error) {
	for _, currentUser := range hu.users {
		if currentUser.Id == userId {
			return &currentUser, nil
		}

	}
	return nil, httperr.ErrNotFound
}

func (hu *HandlerUser) Update(user *User) (*User, error) {
	for index, currentUser := range hu.users {
		if currentUser.Id != user.Id {
			continue
		}
		hu.users[index] = *user
		return user, nil
	}

	return nil, httperr.ErrInternalFailure
}

func (hu *HandlerUser) DeleteById(userId UserId) (*User, error) {
	fmt.Println(userId)
	for index, currentUser := range hu.users {
		if currentUser.Id != userId {
			continue
		}
		hu.users = RemoveElement[User](hu.users, index)
		return &currentUser, nil
	}

	return nil, httperr.ErrNotFound
}
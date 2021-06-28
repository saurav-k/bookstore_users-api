package users

import (
	"fmt"

	"github.com/saurav-k/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFounderror(fmt.Sprintf("User %d not fouund", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadrequest(fmt.Sprintf("email %d is already registered", user.Email))
		}
		return errors.NewBadrequest(fmt.Sprintf("user %d is already present", user.Id))

	}
	userDB[user.Id] = user
	return nil
}

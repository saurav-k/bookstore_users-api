package users

import (
	"fmt"
	"github.com/saurav-k/bookstore_users-api/datasources/mysql/users_db"
	"github.com/saurav-k/bookstore_users-api/utils/date_utils"
	"github.com/saurav-k/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUsers = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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

	//current := userDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadrequest(fmt.Sprintf("email %d is already registered", user.Email))
	//	}
	//	return errors.NewBadrequest(fmt.Sprintf("user %d is already present", user.Id))
	//
	//}
	user.DateCreated = date_utils.GetNowString()

	stmt, err := users_db.Client.Prepare(queryInsertUsers)
	if err != nil {
		return  errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// Create insert statement
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil{
		if strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadrequest(
				fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error while inserting user : %s ", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil{
		return errors.NewInternalServerError(
			fmt.Sprintf("error while inserting user : %s ", err.Error()))
	}

	// Assign the auto incremented user ID
	user.Id = userID
	return nil
}

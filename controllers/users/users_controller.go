package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saurav-k/bookstore_users-api/domain/users"
	"github.com/saurav-k/bookstore_users-api/services"
	"github.com/saurav-k/bookstore_users-api/utils/errors"
)

// Will be defined in Phase 1
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadrequest("invalid json body")
		c.JSON(int(restErr.Status), restErr)
		return
	}
	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(int(saveErr.Status), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userError != nil {
		err := errors.NewBadrequest("Invalid user id")
		c.JSON(int(err.Status), err)
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(int(getErr.Status), getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Will be defined in Phase 2
func GetAllUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}

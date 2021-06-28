package app

import (
	"github.com/saurav-k/bookstore_users-api/controllers/health"
	"github.com/saurav-k/bookstore_users-api/controllers/users"
)

func mapUrl() {
	router.GET("/health", health.Health)
	// Phase 1 API
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users", users.GetAllUsers)

	// Phase 2 API
	// router.GET("/users/search/:user", users.SearchUser)
	// router.PUT("/users", users.UpdateUser)
}

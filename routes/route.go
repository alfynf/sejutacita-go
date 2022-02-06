package routes

import (
	"sejutacita/constant"
	"sejutacita/controllers"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// User Sign Up & Sign In
	e.POST("/login", controllers.LoginUsersController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(m.JWT([]byte(constant.SECRET_JWT)))
	r.POST("/users", controllers.CreateUserController)
	r.GET("/users", controllers.GetUserController)
	r.DELETE("/users", controllers.DeleteUserController)
	r.PUT("/users", controllers.UpdateUserController)

	return e
}

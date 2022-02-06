package controllers

import (
	"net/http"
	"regexp"
	"sejutacita/lib/databases"
	"sejutacita/lib/plugins"
	"sejutacita/lib/responses"
	"sejutacita/middlewares"
	"sejutacita/models"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// controller untuk mendaftarkan nasabah baru
func CreateUserController(c echo.Context) error {
	_, role := middlewares.ExtractTokenUser(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Not allowed to create user, please login as admin"))
	}
	var body models.User
	c.Bind(&body)
	body.Name = strings.TrimSpace(body.Name)
	if body.Name == "" || body.Username == "" || body.Email == "" || body.Password == "" || body.Role == "" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Empty Data"))
	}
	if len(body.Password) < 5 {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Password must more than 5 character"))
	}
	// error buat yang duplicate
	if duplicate, _ := databases.GetUserByEmail(bson.M{"email": body.Email}); duplicate != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Email is already used"))
	}

	body.Password, _ = plugins.Encrypt(body.Password)
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(body.Email))
	if !matched {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Wrong email format"))
	}
	_, err := databases.CreateUser(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Failed to create new data"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponse())
}

// fungsi untuk melihat daftar semua user yang sudah pernah terdaftar
func GetUserController(c echo.Context) error {
	username, role := middlewares.ExtractTokenUser(c)
	if role == "admin" {
		user, err := databases.GetUser(bson.M{})
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.FailedReponse("Failed to get data"))
		}
		return c.JSON(http.StatusOK, responses.SuccessResponseWithData(user))
	}
	user, err := databases.GetUserByUsername(bson.M{"username": username})
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Failed to get data"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseWithData(user))
}

// fungsi untuk melihat daftar semua user yang sudah pernah terdaftar
func DeleteUserController(c echo.Context) error {
	_, role := middlewares.ExtractTokenUser(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Not allowed to delete user, please login as admin"))
	}
	username := c.QueryParam("username")
	_, err := databases.DeleteUser(bson.M{"username": username})
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Failed to delete data"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponse())
}

// fungsi untuk melihat daftar semua user yang sudah pernah terdaftar
func UpdateUserController(c echo.Context) error {
	_, role := middlewares.ExtractTokenUser(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Not allowed to update user, please login as admin"))
	}
	var body models.User
	c.Bind(&body)
	username := c.QueryParam("username")
	body.Name = strings.TrimSpace(body.Name)
	if body.Name == "" || body.Username == "" || body.Email == "" || body.Password == "" || body.Role == "" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Empty Data"))
	}
	if len(body.Password) < 5 {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Password must more than 5 character"))
	}

	body.Password, _ = plugins.Encrypt(body.Password)
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(body.Email))
	if !matched {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Wrong email format"))
	}
	// data, _ := databases.ToDoc(body)
	_, err := databases.UpdateUser(body, bson.M{"username": username})
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Failed to update data"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponse())
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, err := databases.Login(user)
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(user.Email))
	if user.Email == "" && user.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Empty email and password data"))
	} else if !matched {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Wrong email format"))
	} else if err != nil || users == 0 {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Login failed"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseWithData(users))
}

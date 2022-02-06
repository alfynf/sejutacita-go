package middlewares

import (
	"fmt"
	"sejutacita/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
	SigningKey: []byte(constant.SECRET_JWT),
})

func CreateToken(username string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 1 hour

	tokenString, err := token.SignedString([]byte(constant.SECRET_JWT))
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ExtractTokenUser(c echo.Context) (string, string) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		role := claims["role"].(string)
		return username, role
	}
	return "", ""
}

package main

import "net/http"
import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"

//User lkjfhlg
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func grtUsers(c echo.Context) error {
	guy := User{
		"teerapong",
		"tuntawin",
		"Guy",
		"123",
	}
	return c.JSON(http.StatusOK, guy)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", index)
	e.GET("/users", grtUsers)
	e.Logger.Fatal(e.Start(":1323"))
}

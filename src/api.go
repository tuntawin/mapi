package main

import "net/http"
import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"

//User struct
type User struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func getUsers(c echo.Context) error {
	guy := User{
		"teerapong",
		"tuntawin",
		"Guy",
		"",
	}
	return c.JSON(http.StatusOK, guy)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", index)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserByID)
	e.POST("users", saveUser)
	e.Logger.Fatal(e.Start(":1323"))
}

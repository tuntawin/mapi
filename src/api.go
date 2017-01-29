package main

import "net/http"
import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"


func index(c echo.Context) error {
  return c.JSON(http.StatusOK, "Hello World")
}

func grtUsers(c echo.Context) error {
  firstname := "Teerapong"
  return c.JSON(http.StatusOK, firstname)
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.GET("/", index)
  e.GET("/users", grtUsers)
  e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"log"

	"github.com/AnuchitO/demopkg/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	user.InitDB()

	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "apidesign" || password == "45678" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", user.CreateUserHandler)
	e.GET("/users", user.GetUsersHandler)
	e.GET("/users/:id", user.GetUserHandler)

	log.Fatal(e.Start(":2565"))
}

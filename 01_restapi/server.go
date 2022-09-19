package main

import (
	"net/http"
	"os"

	"github.com/AnuchitO/restapi/customer"
	"github.com/labstack/echo/v4"
)

func setupRoute() *echo.Echo {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return next(c)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token != os.Getenv("AUTH_TOKEN") {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	})

	e.POST("/customers", customer.CreateHandler)
	e.GET("/customers/:id", customer.GetCustomerByIdHandler)
	e.GET("/customers", customer.GetCustomersHandler)
	e.PUT("/customers/:id", customer.UpdateCustomerHandler)
	e.DELETE("/customers/:id", customer.DeleteCustomerHandler)

	return e
}

func main() {
	r := setupRoute()
	r.Start(":2565")
}

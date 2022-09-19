package customer

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Err struct {
	Message string `json:"message"`
}

func CreateHandler(c echo.Context) error {
	cust := Customer{}
	err := c.Bind(&cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO customers(name, email, status) VALUES($1, $2, $3) RETURNING id;", cust.Name, cust.Email, cust.Status)
	err = row.Scan(&cust.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	fmt.Printf("id : % #v\n", cust)

	return c.JSON(http.StatusCreated, cust)
}

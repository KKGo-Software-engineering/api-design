package customer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCustomerByIdHandler(c echo.Context) error {
	id := c.Param("id")

	rowID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "id should be int " + err.Error()})
	}

	row := db.QueryRow("SELECT id, name, email, status FROM customers WHERE id=$1", rowID)

	cst := Customer{}
	err = row.Scan(&cst.ID, &cst.Name, &cst.Email, &cst.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	fmt.Printf("cst % #v\n", cst)

	return c.JSON(http.StatusOK, cst)
}

func GetCustomersHandler(c echo.Context) error {
	custs := []Customer{}

	rows, err := db.Query("SELECT id, name, email, status FROM customers")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	for rows.Next() {
		cst := Customer{}
		err := rows.Scan(&cst.ID, &cst.Name, &cst.Email, &cst.Status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})

		}
		custs = append(custs, cst)
	}

	return c.JSON(http.StatusOK, custs)
}

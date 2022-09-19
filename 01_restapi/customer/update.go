package customer

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateCustomerHandler(c echo.Context) error {
	rowID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	cust := Customer{}
	// find by id
	err = c.Bind(&cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	stmt, err := db.Prepare(`
	UPDATE customers
	SET name=$2, email=$3, status=$4
	WHERE id=$1
	`)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	if _, err := stmt.Exec(rowID, cust.Name, cust.Email, cust.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	cust.ID = rowID
	return c.JSON(http.StatusOK, cust)
}

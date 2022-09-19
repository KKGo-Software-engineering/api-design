package customer

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteCustomerHandler(c echo.Context) error {
	id := c.Param("id")

	rowID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	stmt, err := db.Prepare("DELETE FROM customers WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if _, err := stmt.Exec(rowID); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Err{Message: "customer deleted"})
}

package api

import (
	"net/http"
	"github.com/labstack/echo"
)


func HeathCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
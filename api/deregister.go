package api

import (
	"net/http"
	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"

	"github.com/vietwow/getenv/model"
)


func DeRegisterService(c echo.Context) error {
	AppName := c.Param("AppName")

    consul_client.DeleteKeyValue(AppName)

	return c.NoContent(http.StatusNoContent)
}
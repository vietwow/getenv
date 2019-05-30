package api

import (
	"net/http"
	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"

	"github.com/vietwow/getenv/model"
)


func DeRegisterService(c echo.Context) error {
	// Bind the input data to ExampleRequest
	Request := new(model.Request)
	if err := c.Bind(Request); err != nil {
		return err
	}

    //
    consul_client.DeleteKeyValue(Request.AppName)

	return c.NoContent(http.StatusNoContent)
}
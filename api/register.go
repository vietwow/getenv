package api

import (
	"net/http"
	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"

	"github.com/vietwow/getenv/model"
)


func RegisterService(c echo.Context) error {
	// Bind the input data to ExampleRequest
	Request := new(model.RequestRegister)
	if err := c.Bind(Request); err != nil {
		return err
	}

    //
    consul_client.StoreKeyValue(Request.AppName, Request.URL)

	return c.JSON(http.StatusCreated, Request.AppName)
}
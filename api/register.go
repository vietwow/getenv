package api

import (
	"fmt"

	"net/http"
	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"

	"github.com/vietwow/getenv/model"
)


func Register(c echo.Context) {
	// Bind the input data to ExampleRequest
	Request := new(model.RequestRegister)
	if err := c.Bind(Request); err != nil {
		return err
	}

    //
    err := consul_client.RegisterService(Request.AppName, Request.URL)

	return c.JSON(http.StatusCreated, Request.AppName)
}
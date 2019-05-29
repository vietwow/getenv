package api

import (
	"fmt"

	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"

	"github.com/vietwow/getenv/model"
)


func GetURL(c echo.Context) {
	// Bind the input data to ExampleRequest
	Request := new(model.Request)
	if err := c.Bind(Request); err != nil {
		return err
	}

    //
    URL := consul_client.RegisterService(Request.AppName)

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
		  fmt.Sprintf(`{
		    "URL": %s
		  }`, URL),
		),
	)
}
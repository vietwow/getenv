package api

import (
	"time"

	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getevn/pkg/consul_client"

	"gitlab.com/vietwow/getenv/model"
)


func GetURL(c echo.Context) {
	// Bind the input data to ExampleRequest
	Request := new(model.Request)
	if err := c.Bind(Request); err != nil {
		return err
	}

    // Init consul client
    cClient := consul_client.NewConsulClient()

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
		  fmt.Sprintf(`{
		    "URL": %s
		  }`, Request.URL),
		),
	)
}
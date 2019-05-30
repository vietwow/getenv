package api

import (
	"fmt"

	"net/http"
	"github.com/labstack/echo"
	// consul_client "github.com/vietwow/getenv/pkg/consul"

	// "github.com/vietwow/getenv/model"
)


func List(c echo.Context) error {
    // err := consul_client.ListAllService()
    services := "abc"

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
		  fmt.Sprintf(`{
		    "Services": %s
		  }`, services),
		),
	)
}
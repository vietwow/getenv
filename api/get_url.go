package api

import (
	"fmt"

	"net/http"
	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"

	"github.com/vietwow/getenv/model"
)


func GetURL(c echo.Context) error {
	// Bind the input data to ExampleRequest
	Request := new(model.Request)
	if err := c.Bind(Request); err != nil {
		return err
	}

    //
    URL := consul_client.GetURLFromAppName(Request.AppName)
    fmt.Println(URL)
    // URL := "hahaha"

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
		  fmt.Sprintf(`{
		    "URL": %s
		  }`, URL),
		),
	)
}
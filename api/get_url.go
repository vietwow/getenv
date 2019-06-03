package api

import (
	"fmt"

	"net/http"
	"github.com/labstack/echo"
	consul_client "github.com/vietwow/getenv/pkg/consul"
)


func GetURLFromAppName(c echo.Context) error {
    AppName := c.Param("AppName")

    URL := consul_client.GetKeyValue(AppName)

    if URL == "" {
    	return c.NoContent(http.StatusNoContent)
    }

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
		  fmt.Sprintf(`%s`, URL),
		),
	)
}
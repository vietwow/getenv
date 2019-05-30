package main

import (
	"github.com/labstack/echo"
	"github.com/vietwow/getenv/api"
)

func main() {
    e := echo.New()

    // Routes
    // e.GET("/url/:url", api.GetURL())
    // e.GET("/register/:url", api.Register())
    // e.GET("/deregister/:url", api.DeRegister())
    // e.GET("/list/:url", api.List())
    e.GET("/services", api.List)
    e.GET("/services/:AppName", api.GetURL)
    e.POST("/services", api.Register)
    e.PUT("/services", api.Register)
    e.DELETE("/services/:id", api.DeRegister)

    // Start server at localhost:8080
    e.Logger.Fatal(e.Start(":8080"))

}

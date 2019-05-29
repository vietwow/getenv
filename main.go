package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/labstack/echo"
	"github.com/vietwow/getevn/api"
)


func main() {
    kingpin.Version(version)
    kingpin.Parse()

    e := echo.New()

    // Routes
    e.GET("/getenv/:url", api.GetURL())

    // Start server at localhost:8080
    e.Logger.Fatal(e.Start(":8080"))

}

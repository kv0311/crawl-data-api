package main

import (
	"fmt"

	"crawl-project/route"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	route.CrawlRoute(e)
	fmt.Println("Welcome to the webserver")
	e.Start(":3030")
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"crawl-project/route"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	route.CrawlRoute(e)
	fmt.Println("Welcome to the webserver")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" // Default port if not specified
	}
	err := http.ListenAndServe(":"+port, e)
	if err != nil {
		fmt.Println(err)
	}
}

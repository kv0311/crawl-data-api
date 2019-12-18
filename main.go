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
	if len(port) == 0 {
		port = "9999"
	}
	fmt.Println("listening on port: ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

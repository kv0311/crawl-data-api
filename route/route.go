package route

import (
	"crawl-project/handler"
	"github.com/labstack/echo"
)

func CrawlRoute(e *echo.Echo) {
	e.GET("/", handler.CrawlData)
	e.GET("/get-all", handler.MoveDataBase1)
	e.GET("/migrate-database-2", handler.MigrateToDataBase2)
	e.GET("/get/all", handler.GetAll)
	e.GET("/get-by-name", handler.GetByName)
	e.GET("/get-by-phone", handler.GetByPhone)

}
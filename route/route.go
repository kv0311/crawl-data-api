package route

import (
	"crawl-project/handler"
	"github.com/labstack/echo"
)

// CrawlRoute ...
func CrawlRoute(e *echo.Echo) {
	// Api để crawl và migrate data
	// e.GET("/", handler.CrawlData)
	// e.GET("/get-all", handler.MoveDataBase1)
	// e.GET("/migrate-database-2", handler.MigrateToDataBase2)
	e.GET("/get/all/database-1", handler.GetAllData1)
	e.GET("/get/all/database-2", handler.GetAll)
	e.GET("/get-by-name", handler.GetByName)
	e.GET("/get-by-phone", handler.GetByPhone)

}

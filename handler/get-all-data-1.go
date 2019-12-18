package handler

import (
	"crawl-project/model"
	"crawl-project/repo"

	"github.com/labstack/echo"
)

// GetAllData1 ...
func GetAllData1(c echo.Context) (err error) {
	var data []model.DataCrawl
	data, err = repo.GetAllNewDatabase1()
	if err != nil {
		return
	}
	return c.JSON(400, data)
}

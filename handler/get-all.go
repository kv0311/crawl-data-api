package handler

import (
	"crawl-project/model"
	"crawl-project/repo"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) (err error) {
	var data []model.DataCrawl2
	data, err = repo.GetAllNewDatabase2()
	if err != nil {
		return
	}
	return c.JSON(400, data)
}

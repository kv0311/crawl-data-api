package handler

import (
	"crawl-project/model"
	"crawl-project/repo"

	"github.com/labstack/echo"
)

// GetByName ...
func GetByName(c echo.Context) (err error) {
	type myRequest struct {
		Name string `json:"name" query:"name"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	var data []model.DataCrawl2
	data, err = repo.GetDataByName(request.Name)
	if err != nil {
		return
	}
	if len(data) == 0 {
		return c.JSON(404, "Không tìm thấy thông tin")
	}
	return c.JSON(400, data)
}

package handler

import (
	"crawl-project/model"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"crawl-project/repo"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)

//Đây là những api chỉ để crawl data
//CrawlData and insert to database old
func CrawlData(c echo.Context) (err error) {
	var dataArray []model.DataCrawl
	for i := 1; i < 1000; i++ {

		data, _ := CrawlDataFunction(i)
		for k := 0; k < len(data); k++ {
			repo.InsertOneData(data[k])
			time.Sleep(100)
		}
	}
	return c.JSON(400, dataArray)
}

// MigrateToDataBase2 ...
func MigrateToDataBase2(c echo.Context) (err error) {
	dataArray, err := repo.GetAllNewDatabase1()
	if err != nil {
		return
	}
	var data2 model.DataCrawl2
	var data2Array []model.DataCrawl2
	for i := 0; i < len(dataArray); i++ {
		data2.Name = dataArray[i].Name
		data2.Phone = dataArray[i].Phone
		data2.Address = dataArray[i].Address
		data2.Owner = strings.Contains(strings.ToLower(strings.Replace(dataArray[i].Title, " ", "", -1)), "chínhchủ")
		data2Array = append(data2Array, data2)
	}
	arrayData2Interface := []interface{}{}
	for i := 0; i < len(data2Array); i++ {
		data := &data2Array[i]
		var dataInterface interface{}
		marshalValue, _ := json.Marshal(data)
		json.Unmarshal(marshalValue, &dataInterface)
		arrayData2Interface = append(arrayData2Interface, dataInterface)
	}
	repo.InsertToNewDatabase2(arrayData2Interface)
	return c.JSON(400, arrayData2Interface)
}

//MoveDataBase1 from old database
func MoveDataBase1(c echo.Context) (err error) {
	arrayData, _ := repo.GetAllData()

	arrayDataInterface := []interface{}{}
	for i := 0; i < len(arrayData); i++ {
		data := &arrayData[i]
		var dataInterface interface{}
		marshalValue, _ := json.Marshal(data)
		json.Unmarshal(marshalValue, &dataInterface)
		arrayDataInterface = append(arrayDataInterface, dataInterface)
	}
	repo.InsertToNewDatabase(arrayDataInterface)
	return c.JSON(400, arrayDataInterface)
}

//CrawlDataFunction ...
//Crawl data form batdongsan.com.vn
func CrawlDataFunction(a int) (result []model.DataCrawl, err error) {
	url := "https://batdongsan.com.vn/ban-dat/p" + strconv.Itoa(a)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	// s, _ := doc.Find("div.p-main-text").Attr("")
	// g := s.Find("d").Text()
	var titleArray []string
	var contentArray []string
	var imageArray []string
	var priceArray []string
	var areaArray []string
	var districtArray []string
	var uptimeArray []string
	var hrefArray []string
	var data model.DataCrawl
	var dataArray []model.DataCrawl
	doc.Find("div.Main div.search-productItem").Each(func(i int, s *goquery.Selection) {
		// lastPageLink, exists := doc.Find("div.p-title a").Attr("title")
		// if exists != true {
		// 	lastPageLink = "#"
		// }
		// Title := s.Find("div.p-title a").Text()
		// a = append(a, Title)
		//Get title
		title, _ := s.Find("div.p-title a").Attr("title")
		titleArray = append(titleArray, title)
		//Get content
		content := s.Find("div.p-main-text").Text()
		contentArray = append(contentArray, content)
		//Get image
		image, _ := s.Find("img.product-avatar-img").Attr("src")
		imageArray = append(imageArray, image)
		//Get price
		price := s.Find("strong.product-price").Text()
		priceArray = append(priceArray, price)
		//Get area
		area := s.Find("strong.product-area").Text()
		areaArray = append(areaArray, area)
		//Get district
		district := s.Find("strong.product-city-dist").Text()
		districtArray = append(districtArray, district)
		//Get uptime
		uptime := s.Find("span.uptime").Text()
		uptimeArray = append(uptimeArray, uptime)
		//Get href
		href, _ := s.Find("a.product-avatar").Attr("href")
		hrefArray = append(hrefArray, href)
		urlChild := "https://batdongsan.com.vn" + href
		docum, err := goquery.NewDocument(urlChild)
		if err != nil {
			return
		}
		g := docum.Find("div#divCustomerInfo")
		name := g.Find("div#LeftMainContent__productDetail_contactName div.right").Text()
		nameSplit := strings.Replace(name, "\n", "", -1)
		address := g.Find("div#LeftMainContent__productDetail_contactAddress div.right").Text()
		addressSplit := strings.Replace(address, "\n", "", -1)
		phone := g.Find("div#LeftMainContent__productDetail_contactMobile div.right").Text()
		phoneSplit := strings.Replace(phone, "\n", "", -1)
		data.Phone = phoneSplit
		data.Name = nameSplit
		data.Address = addressSplit
		data.Title = title
		data.Content = content
		data.Image = image
		data.Price = price
		data.Area = area
		data.District = district
		data.UpTime = uptime
		dataArray = append(dataArray, data)
	})

	result = dataArray
	return
}

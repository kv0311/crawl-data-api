package model

//DataCrawl ...
type DataCrawl struct {
	Title    string `json:"title" bson:"title"`
	Content  string `json:"content" bson:"content"`
	Image    string `json:"image" bson:"image"`
	Price    string `json:"price" bson:"price"`
	Area     string `json:"area" bson:"area"`
	District string `json:"district" bson:"district"`
	UpTime   string `json:"up_time" bson:"up_time"`
	Name     string `json:"name" bson:"name"`
	Phone    string `json:"phone" bson:"phone"`
	Address  string `json:"address" bson:"address"`
}

// DataCrawl2 ...
type DataCrawl2 struct {
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
	Owner   bool   `json:"owner" bson:"owner"`
}

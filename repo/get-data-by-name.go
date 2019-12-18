package repo

import (
	"context"
	"crawl-project/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDataByName ...
func GetDataByName(name string) (dataArray []model.DataCrawl2, err error) {
	uri := "mongodb+srv://LevineNguyen:Khanhvinh1998@cluster0-2bvwu.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Println(err)
		return
	}
	var cursor *mongo.Cursor
	collection := client.Database("DataCrawl").Collection("database1")
	cursor, err = collection.Find(context.TODO(), bson.M{"name": name})
	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			result := model.DataCrawl2{}
			err = cursor.Decode(&result)
			fmt.Println(result)
			if err != nil {
				return nil, err
			}

			dataArray = append(dataArray, result)
		}

		// Close the cursor once finished
		_ = cursor.Close(context.TODO())
	}
	return dataArray, nil
}

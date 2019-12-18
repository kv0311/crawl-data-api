package repo

import (
	"context"
	"fmt"
	"log"

	"crawl-project/model"

	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

//InsertOneData : insert data to localdb
func InsertOneData(data model.DataCrawl) (err error) {
	uri := "localhost:27017"
	session, err := mgo.Dial(uri)
	// collection := mongoconfig.InitMongo("DataCrawl", "db1")
	collection := session.DB("DataCrawl").C("database1")
	err = collection.Insert(data)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	defer session.Close()
	return
}

//GetAllData get all data from local db
func GetAllData() (dataArray []model.DataCrawl, err error) {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Println(err)
		return
	}
	var cursor *mongo.Cursor
	collection := client.Database("DataCrawl").Collection("database1")
	cursor, err = collection.Find(context.TODO(), bson.M{})
	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			result := model.DataCrawl{}
			err = cursor.Decode(&result)
			if err != nil {
				return nil, err
			}

			dataArray = append(dataArray, result)
		}

		// Close the cursor once finished
		_ = cursor.Close(context.TODO())
	}
	// for a.Iter().Next(context.TODO()) {
	// 	data := model.DataCrawl{}
	// 	fmt.Println(a)
	// 	dataArray = append(dataArray, data)
	// }

	return dataArray, nil

}

//InsertToNewDatabase insert data from local db to database1 (mograte success)
func InsertToNewDatabase(dataInput []interface{}) (dataArray []model.DataCrawl, err error) {
	uri := "mongodb+srv://LevineNguyen:Khanhvinh1998@cluster0-2bvwu.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
		return
	}
	collection := client.Database("DataCrawl").Collection("database1")
	_, err = collection.InsertMany(context.TODO(), dataInput)
	// for a.Iter().Next(context.TODO()) {
	// 	data := model.DataCrawl{}
	// 	fmt.Println(a)
	// 	dataArray = append(dataArray, data)
	// }

	return dataArray, nil
}

// GetAllNewDatabase1 ... get all data from database1
func GetAllNewDatabase1() (dataArray []model.DataCrawl, err error) {
	uri := "mongodb+srv://LevineNguyen:Khanhvinh1998@cluster0-2bvwu.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Println(err)
		return
	}
	var cursor *mongo.Cursor
	collection := client.Database("DataCrawl").Collection("database1")
	cursor, err = collection.Find(context.TODO(), bson.M{})
	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			result := model.DataCrawl{}
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

// InsertToNewDatabase2 : insert data from database 1 to database 2
func InsertToNewDatabase2(dataInput []interface{}) (dataArray []model.DataCrawl, err error) {
	uri := "mongodb+srv://LevineNguyen:Khanhvinh1998@cluster0-2bvwu.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
		return
	}
	collection := client.Database("DataCrawl").Collection("database2")
	_, err = collection.InsertMany(context.TODO(), dataInput)
	if err != nil {
		return
	}
	return dataArray, nil
}

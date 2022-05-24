package models

import (
	"context"
	"item_golang/utils"
	"log"
	"time"

	dataAccessLayer "github.com/ansh1207/dataAccessLayer/db"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoConfigType struct {
	DbType int
	DbUrl  string
	DbName string
}

func GetAppDetailsFromDB(appChan chan<- interface{}) {
	config, _ := utils.LoadConfiguration()

	mongoClient := dataAccessLayer.NewStore(MongoConfigType{DbType: config.DbType, DbUrl: config.DbUrl, DbName: config.DbName})
	mongoClient.Connect()
	defer mongoClient.Cancel()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoData, err := mongoClient.FindOne(ctx, "constants", bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	appChan <- mongoData
}

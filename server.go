package main

import (
	controller "item_golang/controller"
	"item_golang/models"
	utils "item_golang/utils"
	"log"
	"net/http"

	dataAccessLayer "github.com/ansh1207/dataAccessLayer/db"
)

func handleRequests() {
	config, _ := utils.LoadConfiguration()
	http.HandleFunc("/listitem", controller.ListItems)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func main() {
	config, _ := utils.LoadConfiguration()
	MongoClient := dataAccessLayer.NewStore(models.MongoConfigType{DbType: config.DbType, DbUrl: config.DbUrl, DbName: config.DbName})
	MongoClient.Connect()
	handleRequests()
}

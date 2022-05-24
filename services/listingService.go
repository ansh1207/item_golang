package services

import (
	"fmt"
	"item_golang/models"
)

type AppService interface {
	GetAppsDetailsFromCache(packages PackageList, idList IdList, filters []interface{}) []App
	GetUnavailableApps(apps App) []App
	GetAppDetailsFromDB(unavailableApps PackageList, lean bool)
	GetFromCacheBulk(profile models.RedisProfiles, packages PackageList) (App, error)
	SetInCacheBulk(profile models.RedisProfiles, unavailableApps []App)
	PrepareCacheUpdates(appsInDB []App)
}

type PackageList []struct {
	PackageName string
}

type IdList []struct {
	ID string
}

type App struct {
}

// func (app App) GetAppsDetailsFromCache(packages PackageList, idList IdList, filters []interface{}) []App {
// 	// return
// }

// func (app App) GetUnavailableApps(apps App) []App {
// 	return
// }

// func (app App) GetAppDetailsFromDB

func FindItem() {

	// var availableApps, unavailableApps, appsFoundInDb, _apps = []App{}, []App{}, []App{}, []App{}

	// go models.GetAppsDetailsFromCache()

	// go GetUnavailableApps()

	applicationChan := make(chan interface{}, 1)

	go models.GetAppDetailsFromDB(applicationChan)

	appData := <-applicationChan

	fmt.Println(appData)

}

// func GetUnavailableApps() {

// }

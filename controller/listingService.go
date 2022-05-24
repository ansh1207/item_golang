package controller

import (
	"encoding/json"
	services "item_golang/services"
	utils "item_golang/utils"
	"net/http"
)

type listingService interface {
	fetchData(w http.ResponseWriter, r *http.Request) (services.App, error)
}

type AppRequest struct {
}

type AppStoriesRequest struct {
}

type Filters struct {
}

func (app AppRequest) fetchData(w http.ResponseWriter, r *http.Request) (services.App, error) {
	defer utils.Elapsed("Fetching data")()

	services.FindItem()

	return services.App{}, nil
}

func ListItems(w http.ResponseWriter, r *http.Request) {
	var itemList listingService = AppRequest{}

	res, err := itemList.fetchData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resBody, _ := json.Marshal(res)

	w.Write([]byte(resBody))

}

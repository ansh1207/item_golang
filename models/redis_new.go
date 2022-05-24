package models

type RedisProfiles struct {
	AppInfo []struct {
		Type   string `json:"type"`
		Prefix string `json:"prefix"`
		Ttl    int    `json:"ttl"`
	}
}

func GetAppsDetailsFromCache() {

}

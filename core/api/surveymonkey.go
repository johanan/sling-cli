package api

import (
	"fmt"

	h "github.com/flarco/gutil"
)

// SurveyMonkeyAPI is for surveymonkey
// https://developer.surveymonkey.com/api/v3
type SurveyMonkeyAPI struct {
	BaseAPI
}

// Init initializes
func (api *SurveyMonkeyAPI) Init() (err error) {
	api.Provider = SurveyMonkey
	api.BaseURL = "https://api.surveymonkey.com/v3"
	api.Key = api.properties["SURVEYMONKEY_ACCESS_TOKEN"]

	if api.Key == "" {
		err = h.Error(fmt.Errorf("did not provide SURVEYMONKEY_ACCESS_TOKEN"))
		return
	}

	api.DefHeaders = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "bearer " + api.Key,
	}

	return api.BaseAPI.Init()
}

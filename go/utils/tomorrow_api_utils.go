package utils

import (
	"fmt"
	"reflect"
	"smart_vibes/models"
)

const BaseUrl string = "https://api.tomorrow.io/v4/"

// EncodeTomorrowUrl takes path or the route with parameters as map which then gets converted to a single url
// combined with base url of Tomorrow.io api
func EncodeTomorrowUrl(path string, parameters map[string]any) string {

	var parametersAsString = "?"
	for k, v := range parameters {
		valueType := reflect.TypeOf(v).Kind()

		if k == "timesteps" {
			encodeTimeSteps(&parametersAsString, v.([]models.Timestep))
		} else {

			if valueType == reflect.String {
				parametersAsString += k + "=" + v.(string) + "&"
			}

		}

	}
	parametersAsString += "apikey=" + ApiKey
	encodedUrl := BaseUrl + path + parametersAsString
	return encodedUrl

}

func encodeTimeSteps(paramsAsString *string, params []models.Timestep) {
	for _, v := range params {
		*paramsAsString += fmt.Sprintf("timesteps=%v&", v)
	}
}

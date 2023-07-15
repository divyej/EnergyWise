package utils

import (
	"fmt"
	"reflect"
)

const BaseUrl string = "https://api.tomorrow.io/v4/"

func EncodeParameters(path string, parameters map[string]any) string {

	var parametersAsString = ""
	for k, v := range parameters {
		valueType := reflect.TypeOf(v).Kind()

		if k == "timesteps" {
			encodeTimeSteps(&parametersAsString, v.([]string))
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

func encodeTimeSteps(paramsAsString *string, params []string) {
	for _, v := range params {
		*paramsAsString += fmt.Sprintf("timesteps=%v&", v)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/shiroyasha/semaphore/config"
)

func ApiPostRequest(config *config.Config, path string, params map[string]string) ([]byte, error) {
	key_value_pairs := "auth_token=" + config.ApiToken

	if params != nil {
		for key, value := range params {
			key_value_pairs += fmt.Sprintf("&%s=%s", key, value)
		}
	}

	api_url := fmt.Sprintf("%s%s", config.ApiDomain, path)

	form_values, err := url.ParseQuery(key_value_pairs)

	response, err := http.PostForm(api_url, form_values)

	check(err, "Can't connect to Semaphore's API")

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func ApiGetRequest(config *config.Config, path string, params map[string]string) ([]byte, error) {
	key_value_pairs := ""

	if params != nil {
		for key, value := range params {
			key_value_pairs += fmt.Sprintf("&%s=%s", key, value)
		}
	}

	url := fmt.Sprintf("%s%s?auth_token=%s%s", config.ApiDomain, path, config.ApiToken, key_value_pairs)

	response, err := http.Get(url)

	check(err, "Can't connect to Semaphore's API")

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func GetProjects(config *config.Config) Projects {
	content, err := ApiGetRequest(config, "/api/v1/projects", nil)

	check(err, "Can't load the projects from Semaphore")

	var projects Projects

	err = json.Unmarshal(content, &projects)

	check(err, "Can't parse the returned projects JSON")

	return projects
}

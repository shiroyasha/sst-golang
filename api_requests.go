package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ApiPostRequest(path string, params map[string]string) ([]byte, error) {
	token := LoadToken()
	domain := LoadDomain()

	key_value_pairs := "auth_token=" + token

	if params != nil {
		for key, value := range params {
			key_value_pairs += fmt.Sprintf("&%s=%s", key, value)
		}
	}

	api_url := fmt.Sprintf("%s%s", domain, path)

	form_values, err := url.ParseQuery(key_value_pairs)

	response, err := http.PostForm(api_url, form_values)

	check(err, "Can't connect to Semaphore's API")

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func ApiGetRequest(path string, params map[string]string) ([]byte, error) {
	token := LoadToken()
	domain := LoadDomain()

	key_value_pairs := ""

	if params != nil {
		for key, value := range params {
			key_value_pairs += fmt.Sprintf("&%s=%s", key, value)
		}
	}

	url := fmt.Sprintf("%s%s?auth_token=%s%s", domain, path, token, key_value_pairs)

	response, err := http.Get(url)

	check(err, "Can't connect to Semaphore's API")

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func GetProjects() Projects {
	content, err := ApiGetRequest("/api/v1/projects", nil)

	check(err, "Can't load the projects from Semaphore")

	var projects Projects

	err = json.Unmarshal(content, &projects)

	check(err, "Can't parse the returned projects JSON")

	return projects
}

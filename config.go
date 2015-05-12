package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"
)

type SemaphoreConfig struct {
	ApiToken  string
	ApiDomain string
}

func config_cant_load_user() {
	fmt.Println("Fatal error: Can't determine current user")
}

func config_cant_load_api_token() {
	fmt.Println("Can't load api token from the configuration files\n")

	fmt.Println("Please store your Semaphore api_token to ~/.sst/api_token")
}

func config_cant_load_api_domain() {
	fmt.Println("Can't load api domain from the configuration files\n")

	fmt.Println("Please store the Semaphore domain you want ot use to ~/.sst/api_domain")
}

func LoadConfig() (*SemaphoreConfig, error) {
	user, err := user.Current()

	if err != nil {
		config_cant_load_user()
		return nil, err
	}

	api_token, token_err := ioutil.ReadFile(fmt.Sprintf("%s/.sst/api_token", user.HomeDir))
	api_domain, domain_err := ioutil.ReadFile(fmt.Sprintf("%s/.sst/api_domain", user.HomeDir))

	if token_err != nil {
		config_cant_load_api_token()
		return nil, token_err
	}

	if domain_err != nil {
		config_cant_load_api_domain()
		return nil, domain_err
	}

	config := &SemaphoreConfig{
		ApiToken:  strings.TrimSpace(string(api_token)),
		ApiDomain: strings.TrimSpace(string(api_domain)),
	}

	return config, nil
}

package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"
)

func LoadToken() string {
	user, err := user.Current()

	check(err, "Can't load the current user")

	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/.sst/api_token", user.HomeDir))

	check(err, "Can't load the ~/.sst/api_token file")

	return strings.TrimSpace(string(dat))
}

func LoadDomain() string {
	user, err := user.Current()

	check(err, "Can't load the current user")

	domain, err := ioutil.ReadFile(fmt.Sprintf("%s/.sst/api_domain", user.HomeDir))

	if err == nil {
		return strings.TrimSpace(string(domain))
	} else {
		return "" //https://semaphoreci.com"
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"
)

type Branch struct {
	Name   string `json:"branch_name"`
	Result string
}

type Project struct {
	Name     string
	Branches []Branch
}

// Checks if the error is present, and if it is
// this function displays the provided errorMessage and panics
func check(e error, errorMessage string) {
	if e != nil {
		fmt.Printf("%s\n", errorMessage)

		panic(e)
	}
}

// Loads api token from the ~/.sst/api_token file
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
		return "https://semaphoreci.com"
	}
}

// Makes a GET request towards SemaphoreCI's API endpoint,
// collects the projects that the user has can access.
// The returned JSON is parsed into a suitable Go Structure.
func GetProjects(domain, token string) []Project {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/projects?auth_token=%s", domain, token))

	check(err, "Can't load the projects from Semaphore")

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	check(err, "Can't load the projects from Semaphore")

	var projects []Project

	err = json.Unmarshal(content, &projects)

	check(err, "Can't parse the returned projects JSON")

	return projects
}

func BranchColor(b Branch) string {
	switch b.Result {
	case "passed":
		return "\033[32m" // green
	case "failed", "stopped":
		return "\033[31m" // red
	case "pending":
		return "\033[33m" // blue
	default:
		return "\033[34m" // yellow
	}
}

// Draws a project tree
func DrawProjectTree(p Project) {
	fmt.Printf("┌─ %s\n", p.Name)

	for index, b := range p.Branches {
		tree_element := "├──"
		color := BranchColor(b)

		if index == len(p.Branches)-1 {
			tree_element = "└──"
		}

		result := b.Result

		if result == "" {
			result = "unknown"
		}

		fmt.Printf("%s %s%s\033[0m :: %s\n", tree_element, color, result, b.Name)
	}

	fmt.Printf("\n")
}

func main() {
	token := LoadToken()
	domain := LoadDomain()

	projects := GetProjects(domain, token)

	for _, p := range projects {
		DrawProjectTree(p)
	}
}

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
		panic(errorMessage)
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

// Makes a GET request towards SemaphoreCI's API endpoint,
// collects the projects that the user has can access.
// The returned JSON is parsed into a suitable Go Structure.
func GetProjects(token string) []Project {
	response, err := http.Get(fmt.Sprintf("https://s3.semaphoreci.com/api/v1/projects?auth_token=%s", token))

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

	projects := GetProjects(token)

	for _, p := range projects {
		DrawProjectTree(p)
	}
}

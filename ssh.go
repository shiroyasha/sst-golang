package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func LoadSourceJobHashId(build_url string) string {
	page, err := ApiGetRequest(build_url, nil)

	check(err, "Can't load build page")

	r, _ := regexp.Compile("source-job=\"([^\"]*)\"")

	matches := r.FindStringSubmatch(string(page))

	return matches[len(matches)-1]
}

func StartSshSession(source_job_hash_id string) {
	_, err := ApiPostRequest("/ssh_sessions", map[string]string{"job_hash_id": source_job_hash_id})

	check(err, "Can't start ssh session")
}

func RunSshCommand(command string) {
	command_fields := strings.Fields(command)

	cmd := exec.Command(command_fields[0], command_fields[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func Ssh(branch Branch) {
	build_url, err := url.Parse(branch.BuildUrl)

	check(err, "Can't load build url")

	job := LoadSourceJobHashId(build_url.Path)

	StartSshSession(job)

	for i := 0; i < 30; i++ {
		page, err := ApiGetRequest(build_url.Path, nil)

		check(err, "Can't get ssh session's status")

		r, _ := regexp.Compile("ssh-command=\"([^\"]*)\"")

		matches := r.FindStringSubmatch(string(page))

		if len(matches) == 0 {
			fmt.Printf(".")

			time.Sleep(5 * time.Second)
		} else {
			fmt.Printf("\n")

			RunSshCommand(matches[1])

			fmt.Printf("\n")
			break
		}
	}
}

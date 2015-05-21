package main

import (
	"fmt"
	"os"

	"github.com/shiroyasha/semaphore/config"
)

func cli_ssh(config *config.Config) {
	projects := GetProjects(config)

	project_name := os.Args[2]
	branch_name := os.Args[3]

	p, b := projects.FindProjectAndBranch(project_name, branch_name)

	if p == nil {
		fmt.Printf("Can't find project named %s", project_name)
	} else if b == nil {
		fmt.Printf("Can't find branch named %s for project %s", branch_name, project_name)
	} else {
		fmt.Printf("Starting SSH session for %s on %s branch\n", p.Name, b.Name)
		Ssh(config, *b)
	}
}

func cli_status(config *config.Config) {
	projects := GetProjects(config)

	DrawProjects(projects)
}

func cli_usage() {
	fmt.Println()
	fmt.Printf("Usage:")

	fmt.Println()
	fmt.Println()

	fmt.Printf("  semaphore status                            -- show the current status of your projects on Semaphore\n")
	fmt.Printf("  semaphore ssh <project_name> <branch_name>  -- starts an SSH session for the provided project/branch\n")
	fmt.Printf("  semaphore config api.token <api_token>      -- saves the api token to the configuration file\n")
	fmt.Printf("  semaphore config api.domain <api_domain>    -- saves the api domain to the configuration file\n")
	fmt.Printf("  semaphore help                              -- shows this message\n")

	fmt.Println()
}

func cli_usage_error(message string) {
	fmt.Println(message)

	cli_usage()
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "config" {
		if len(os.Args) == 2 {
			cli_usage_error("Field name missing.")
		} else if len(os.Args) == 4 && os.Args[2] == "api.token" {
			config.SaveToken(os.Args[3])
		} else if len(os.Args) == 4 && os.Args[2] == "api.domain" {
			config.SaveDomain(os.Args[3])
		} else {
			cli_usage_error("Unrecognized parameters for config action.")
		}

		return
	}

	c, err := config.Load()

	if err != nil {
		fmt.Println("The api token is not set. Set it with:")
		fmt.Println()
		fmt.Println("  semaphore config api.token <api_token>")
		fmt.Println()

		return
	}

	if len(os.Args) < 2 {
		cli_usage_error("No action specified.")

		return
	}

	switch os.Args[1] {
	case "status":
		if len(os.Args) != 2 {
			cli_usage_error("Status action takes no arguments.")
		} else {
			cli_status(c)
		}
	case "ssh":
		if len(os.Args) != 4 {
			cli_usage_error("SSH action takes 2 arguments")
		} else {
			cli_ssh(c)
		}
	case "help":
		cli_usage()
	default:
		cli_usage_error(fmt.Sprintf("Unrecognized action '%s'.\n", os.Args[1]))
	}
}

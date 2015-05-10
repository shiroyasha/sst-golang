package main

import (
	"fmt"
	"os"
)

func cli_ssh() {
	if len(os.Args) != 4 {
		fmt.Printf("No project and branch options provided\n\n")

		cli_help()
		return
	}

	projects := GetProjects()

	project_name := os.Args[2]
	branch_name := os.Args[3]

	p, b := projects.FindProjectAndBranch(project_name, branch_name)

	if p == nil {
		fmt.Printf("Can't find project named %s", project_name)
	} else if b == nil {
		fmt.Printf("Can't find branch named %s for project %s", branch_name, project_name)
	} else {
		fmt.Printf("Starting SSH session for %s on %s branch\n", p.Name, b.Name)
		Ssh(*b)
	}
}

func cli_status() {
	if len(os.Args) != 2 {
		fmt.Printf("Option <status> takes no arguments\n\n")

		cli_help()

		return
	}

	projects := GetProjects()

	DrawProjects(projects)
}

func cli_help() {
	fmt.Printf("  status                            -- show the current status of your projects on Semaphore\n")
	fmt.Printf("  ssh <project_name> <branch_name>  -- starts an SSH session for the provided project/branch\n")
	fmt.Printf("  help                              -- shows this message\n")
}

func cli_noargs() {
	fmt.Printf("No options provided. Please choose one of the following:\n\n")

	cli_help()
}

func cli_unrecognized() {
	fmt.Printf("Unrecognized option %s\n\n", os.Args[1])

	cli_help()
}

func main() {
	if len(os.Args) < 2 {
		cli_noargs()
		return
	}

	switch os.Args[1] {
	case "status":
		cli_status()
		break
	case "ssh":
		cli_ssh()
		break
	case "help":
		cli_help()
		break
	default:
		cli_unrecognized()
	}
}

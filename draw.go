package main

import "fmt"

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

func DrawBranch(b Branch, last bool) {
	var tree_element string

	if last {
		tree_element = "└──"
	} else {
		tree_element = "├──"
	}

	color := BranchColor(b)

	result := b.Result

	if result == "" {
		result = "unknown"
	}

	branch_status := fmt.Sprintf("%s%s\033[0m", color, result)

	fmt.Printf("%s %s :: %s (%s)\n", tree_element, branch_status, b.Name, b.BuildUrl)
}

func DrawProject(p Project) {
	fmt.Printf("┌─ %s\n", p.Name)

	for index, b := range p.Branches {
		DrawBranch(b, index == len(p.Branches)-1)
	}
}

func DrawProjects(projects []Project) {
	for _, p := range projects {
		DrawProject(p)

		fmt.Printf("\n")
	}
}

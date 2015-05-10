package main

type Projects []Project
type Branches []Branch

type Branch struct {
	Name     string `json:"branch_name"`
	Result   string
	BuildUrl string `json:"build_url"`
}

type Project struct {
	Name     string
	Branches Branches
}

func (projects Projects) FindByName(name string) *Project {
	var result *Project

	for _, project := range projects {
		if project.Name == name {
			result = &project
			break
		}
	}

	return result
}

func (projects Projects) FindProjectAndBranch(project_name, branch_name string) (*Project, *Branch) {
	p := projects.FindByName(project_name)

	if p == nil {
		return nil, nil
	}

	b := p.Branches.FindByName(branch_name)

	return p, b
}

func (branches Branches) FindByName(name string) *Branch {
	var result *Branch

	for _, branch := range branches {
		if branch.Name == name {
			result = &branch
			break
		}
	}

	return result
}

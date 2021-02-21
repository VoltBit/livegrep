package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Failed: %v", err)
	}
}

func makeClient() (*gitlab.Client, error) {
	token := os.Getenv("GITLAB_TOKEN")
	fmt.Println("Token:", token)
	git, err := gitlab.NewClient(token)
	if err != nil {
		log.Fatal("Could not create GitLab client", err)
	}
	return git, err
}

func testUsers() error {
	git, err := makeClient()
	if err != nil {
		return err
	}
	users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	for _, user := range users {
		fmt.Printf("%v\n", user.Name)
	}
	return nil
}

func testProjects() error {
	git, err := makeClient()
	if err != nil {
		return err
	}
	listProjectOpt := gitlab.ListProjectsOptions{
		Visibility: gitlab.Visibility(gitlab.PrivateVisibility),
	}
	// projectList, res, err := git.Projects.ListProjects(&listOptions)
	users, res, err := git.Users.ListUsers(&gitlab.ListUsersOptions{Search: gitlab.String("andreidobre-web")})
	fmt.Printf("Get users response: %v\n", res)
	if err != nil {
		fmt.Printf("Error getting the users: %v", err)
		return err
	}

	var uid int
	for _, user := range users {
		fmt.Printf("User: %v\n", user)
		if user.Username == "andreidobre-web" {
			uid = user.ID
			fmt.Printf(">>> Found user ID: %d\n", uid)
			break
		}
	}

	projectList, res, err := git.Projects.ListUserProjects(uid, &listProjectOpt)
	fmt.Printf("Get projects response: %v\n", res)
	if err != nil {
		return err
	}

	for _, project := range projectList {
		y, m, d := project.CreatedAt.Date()
		fmt.Printf("[%d/%d/%d][%v]: %v\n", d, m, y, project.Owner.Username, project.NameWithNamespace)
	}

	groups, res, err := git.Groups.ListGroups(&gitlab.ListGroupsOptions{})
	fmt.Printf("Get groups response: %v\n%v\n", res, groups)
	if err != nil {
		fmt.Printf("Failed to get groups %v\n", err)
		return err
	}
	for _, group := range groups {
		projs, _, err := git.Groups.ListGroupProjects(group.ID, &gitlab.ListGroupProjectsOptions{})
		if err != nil {
			fmt.Printf("Could not get projects of group: %v\n", err)
		}
		fmt.Printf("Group: %v %v\n", group.FullName, group.FullPath)
		fmt.Println("Group projects:")
		for _, proj := range projs {
			y, m, d := proj.CreatedAt.Date()
			fmt.Printf("[%d/%d/%d][%v]: %v\n", d, m, y, proj.Name, proj.ID)
			files, _, err := git.Repositories.ListTree(proj.ID, &gitlab.ListTreeOptions{})
			if err != nil {
				fmt.Printf("Failed to get the file tree: %v\n", err)
			}
			for _, file := range files {
				fmt.Printf("\t%v: %v\n", file.Path, file.Name)
			}
		}
	}
	return nil
}

func main() {
	checkErr(testProjects())
	// checkErr(testUsers())
}

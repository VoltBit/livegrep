package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

func main() {
	token := os.Getenv("GITLAB_TOKEN")
	fmt.Println("Token:", token)
	git, err := gitlab.NewClient(token)
	if err != nil {
		log.Fatal("Could not create GitLab client", err)
	}
	fmt.Println("Projects:", git.Projects.ListProjects())
}

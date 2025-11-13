package main

import (
	"fmt"
	"net/http"
)

type tempInput struct {
	repositories   []string // input the repositories you want additionally witness
	initRepoNumber int      // entrypoint argument given by ci-cd pipeline
	initRepoTitle  string   // entrypoint argument given by ci-cd pipeline
}

func FetchGithubCICDEnvironment(argument string) {
	resp, err := http.Get(argument)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Kaboom")
	}
}

func AllRepositoryRequests(input tempInput) {

	url := "https://api.github.com/repos/" + input.repositories[1] + "/issues?state=all"
	resp, err := http.Get(url)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Kaboom")
	}

}

func main() {
	exampleInput := tempInput{
		repositories:   []string{"test", "test2"},
		initRepoNumber: 2,
		initRepoTitle:  "idk",
	}
	FetchGithubCICDEnvironment("https://api.github.com/repos/l3montree-dev/devguard-web/pulls/581")
	AllRepositoryRequests(exampleInput)
}

// curl -s https://api.github.com/repos/l3montree-dev/devguard-web/pulls/581
// curl https://api.github.com/repos/l3montree-dev/devguard-documentation/issues?state=all

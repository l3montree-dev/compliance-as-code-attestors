package main

import (
	"fmt"
	"io"
	"net/http"
)

type tempInput struct {
	repositories   []string // input the repositories you want additionally witness
	initRepoNumber int      // entrypoint argument given by ci-cd pipeline
	initRepoTitle  string   // entrypoint argument given by ci-cd pipeline
}

func AllRepositoryRequests(input tempInput) {

	for index, _ := range input.repositories {
		url := "https://api.github.com/repos/" + input.repositories[index] + "/issues?state=all"
		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("whoops")
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		fmt.Printf("%s", body)
		if err != nil {
			fmt.Println("Kaboom")
		}

	}

}

func main() {
	exampleInput := tempInput{
		repositories:   []string{"l3montree-dev/devguard-web", "l3montree-dev/devguard-documentation"},
		initRepoNumber: 581,
		initRepoTitle:  "1277 organization wide dependency search",
	}

	AllRepositoryRequests(exampleInput)
}

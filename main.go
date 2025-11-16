package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type tempInput struct {
	repositories   []string // input the repositories you want additionally witness
	initRepoNumber int      // entrypoint argument given by ci-cd pipeline
	initRepoTitle  string   // entrypoint argument given by ci-cd pipeline
}

type issueSummary struct {
	Repository string `json:"repository"`
	Number     int    `json:"number"`
	Title      string `json:"title"`
}

type combinedOutput struct {
	Repository  string          `json:"repository"`
	IssueNumber int             `json:"issue_number"`
	IssueTitle  string          `json:"issue_title"`
	PullRequest json.RawMessage `json:"pull_request"`
}

var repoIssues []struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func AllRepositoryPRrequests(input tempInput) ([]issueSummary, error) {
	var summaries []issueSummary

	for _, repo := range input.repositories {
		url := "https://api.github.com/repos/" + repo + "/issues?state=all"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("whoops")
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Kaboom")
		}

		if err := json.Unmarshal(body, &repoIssues); err != nil {
			return nil, fmt.Errorf("Kaboom")
		}

		for _, issue := range repoIssues {
			summaries = append(summaries, issueSummary{
				Repository: repo,
				Number:     issue.Number,
				Title:      issue.Title,
			})
		}

	}
	return summaries, nil
}

func AssociatedPullRequest(exampleInput tempInput) {

	summaries, err := AllRepositoryPRrequests(exampleInput)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range summaries {

		// fmt.Printf("general : %s %d: %s\n", s.Repository, s.Number, s.Title)

		if s.Title == exampleInput.initRepoTitle {
			// fmt.Printf("output : %s %d: %s\n", s.Repository, s.Number, s.Title)

			url := "https://api.github.com/repos/" + s.Repository + "/pulls/" + strconv.Itoa(s.Number)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("whoops")
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				fmt.Println("Kaboom")
			}

			output := combinedOutput{
				Repository:  s.Repository,
				IssueNumber: s.Number,
				IssueTitle:  s.Title,
				PullRequest: body,
			}

			finalJSON, err := json.MarshalIndent(output, "", "  ")
			if err != nil {
				log.Printf("failed to marshal response: %v", err)
				continue
			}

			fmt.Printf("%s\n", finalJSON)
		}

		// if err := json.Unmarshal(body, &repoIssues); err != nil {
		// 	return nil, fmt.Errorf("Kaboom")
		// }
	}

}

func main() {
	exampleInput := tempInput{
		repositories:   []string{"l3montree-dev/devguard", "l3montree-dev/devguard-web", "l3montree-dev/devguard-documentation"},
		initRepoNumber: 581,
		initRepoTitle:  "1277 organization wide dependency search",
	}
	AssociatedPullRequest(exampleInput)
}

// curl -s https://api.github.com/repos/l3montree-dev/devguard-web/pulls/581
// curl https://api.github.com/repos/l3montree-dev/devguard-documentation/issues?state=all

// Copyright 2025 larshermges
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cross_pr_request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		repoIssues = nil
		url := "https://api.github.com/repos/" + repo + "/issues?state=all"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("URL GET ERROR")
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Println("Body error")
		}

		if err := json.Unmarshal(body, &repoIssues); err != nil {
			return nil, fmt.Errorf("Unmarshal")
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
	var outputs []combinedOutput

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
				fmt.Printf("PULL ERROR")
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				fmt.Println("reading pull request body error")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("pull request request returned %s: %s\n", resp.Status, strings.TrimSpace(string(body)))
				continue
			}

			output := combinedOutput{
				Repository:  s.Repository,
				IssueNumber: s.Number,
				IssueTitle:  s.Title,
				PullRequest: body,
			}

			outputs = append(outputs, output)
		}

		// if err := json.Unmarshal(body, &repoIssues); err != nil {
		// 	return nil, fmt.Errorf("Kaboom") q
		// }
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(outputs); err != nil {
		log.Printf("failed to marshal response: %v", err)
	}
}

func CrossPRrequest(repos []string, initRepoNumber int, initRepoTitle string) {
	exampleInput := tempInput{
		repositories:   repos,
		initRepoNumber: initRepoNumber,
		initRepoTitle:  initRepoTitle,
	}
	AssociatedPullRequest(exampleInput)
}

// curl -s https://api.github.com/repos/l3montree-dev/devguard-web/pulls/581
// curl https://api.github.com/repos/l3montree-dev/devguard-documentation/issues?state=all

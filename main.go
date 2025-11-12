package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchGithubCICDEnvironment() error {
	const url = "https://api.github.com/repos/l3montree-dev/devguard-web/pulls/581"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	fmt.Printf("GET %s\nStatus: %s\n\n%s\n", url, resp.Status, body)
	return nil
}

func main() {
	if err := FetchGithubCICDEnvironment(); err != nil {
		log.Fatal(err)
	}
}

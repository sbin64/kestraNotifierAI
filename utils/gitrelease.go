package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gofr.dev/pkg/gofr"
)

type GitHubRelease struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	Body        string `json:"body"`
	PublishedAt string `json:"published_at"`
}

func FetchkestraUpdates(c *gofr.Context) (string, error) {
	// Define the GitHub API URL for Kestra releases
	apiURL := "https://api.github.com/repos/kestra-io/kestra/releases"

	// Create a new HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch releases: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var releases []GitHubRelease
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Ensure there are releases available
	if len(releases) == 0 {
		return "No releases found for Kestra.", nil
	}

	// Fetch details of the kestra latest release
	kestra := releases[0]
	kestraContent := fmt.Sprintf(
		"kestra Content: %s\nTag: %s\nDetails:\n%s",
		kestra.Name, kestra.TagName, kestra.Body,
	)

	return kestraContent, nil
}

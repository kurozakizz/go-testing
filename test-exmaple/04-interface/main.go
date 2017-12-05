package main

import (
	"fmt"
	"os"
)

func getReleaseTagMessage(ri ReleaseInfoer, repo string) (string, error) {
	tag, err := ri.GetLatestReleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error querying GitHub API: %s", err)
	}

	return fmt.Sprintf("The latest release is %s", tag), nil
}

func main() {
	gh := GithubReleaseInfo{}
	msg, err := getReleaseTagMessage(gh, "docker/machine")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(msg)
}

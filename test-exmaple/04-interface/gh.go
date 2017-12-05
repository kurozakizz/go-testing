package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubReleaseInfo struct{}

type ReleaseInfo struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

func (gh GithubReleaseInfo) GetLatestReleaseTag(repo string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	response, err := http.Get(apiUrl)
	if err != nil {

		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	releases := []ReleaseInfo{}

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	tag := releases[0].TagName

	return tag, nil
}


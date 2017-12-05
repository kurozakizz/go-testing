package main

import (
	"testing"
)

type StubReleaseInfo struct {
	Tag string
	Err error
}

func (r StubReleaseInfo) GetLatestReleaseTag(repo string) (string, error) {
	if r.Err != nil {
		return "", r.Err
	}

	return r.Tag, nil
}

func TestGetReleaseTagMessage_V010_ShouldReturn_V010(t *testing.T) {
	releaseInfo := StubReleaseInfo{
		Tag: "v0.1.0",
		Err: nil,
	}

	expectedMessage := "The latest release is v0.1.0"
	actualMessage, err := getReleaseTagMessage(releaseInfo, "dev/null")
	if err != nil {
		t.Fatalf("Expected err to be null but it was %s", err)
	}

	if expectedMessage != actualMessage {
		t.Fatalf("Expected %s but got %s", expectedMessage, actualMessage)
	}
}

package main

type ReleaseInfoer interface {
	GetLatestReleaseTag(string) (string, error)
}

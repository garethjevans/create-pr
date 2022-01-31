package cmd

import (
	"fmt"
	"net/http"
)

var _ GitHub = (*DefaultGitHub)(nil)

type GitHub interface {
	PullRequestForBranch(host string, org string, repo string, branch string) (bool, error)
}

type DefaultGitHub struct {
}

func (d DefaultGitHub) PullRequestForBranch(host string, org string, repo string, branch string) (bool, error) {
	// FIXME need to implement host correctly
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?head=%s:%s", org, repo, org, branch)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return true, nil
}

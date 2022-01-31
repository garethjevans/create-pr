package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var _ GitHub = (*DefaultGitHub)(nil)

type HttpGetter interface {
	Get(url string) (*http.Response, error)
}

var Getter = getter

type GitHub interface {
	PullRequestForBranch(host string, org string, repo string, branch string) (bool, error)
	DefaultBranch(host string, org string, repo string) (string, error)
}

type DefaultGitHub struct {
}

type Repository struct {
	DefaultBranch string `json:"default_branch"`
}

func (d DefaultGitHub) DefaultBranch(host string, org string, repo string) (string, error) {
	// FIXME need to implement host correctly
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", org, repo)
	resp, err := Getter().Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	var r Repository
	err = json.Unmarshal(b, &r)
	return r.DefaultBranch, nil
}

func (d DefaultGitHub) PullRequestForBranch(host string, org string, repo string, branch string) (bool, error) {
	// FIXME need to implement host correctly
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?head=%s:%s", org, repo, org, branch)
	resp, err := Getter().Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	var prs []PullRequest
	err = json.Unmarshal(b, &prs)
	return len(prs) > 0, nil
}

type PullRequest struct {
	Number int `json:"number"`
}

func getter() HttpGetter {
	return &http.Client{}
}

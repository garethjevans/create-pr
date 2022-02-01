package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . GitHub

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
	url := fmt.Sprintf("https://%s/repos/%s/%s", determineHost(host), org, repo)
	resp, err := Getter().Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var r Repository
	err = json.Unmarshal(b, &r)
	if err != nil {
		return "", err
	}
	return r.DefaultBranch, nil
}

func (d DefaultGitHub) PullRequestForBranch(host string, org string, repo string, branch string) (bool, error) {
	url := fmt.Sprintf("https://%s/repos/%s/%s/pulls?head=%s:%s", determineHost(host), org, repo, org, branch)
	resp, err := Getter().Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var prs []PullRequest
	err = json.Unmarshal(b, &prs)
	if err != nil {
		return false, err
	}
	return len(prs) > 0, nil
}

type PullRequest struct {
	Number int `json:"number"`
}

func getter() HttpGetter {
	return &http.Client{}
}

func determineHost(h string) string {
	if h == "github.com" {
		return "api.github.com"
	}
	return fmt.Sprintf("%s/api/v3", h)
}

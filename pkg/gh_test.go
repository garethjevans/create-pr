package pkg

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestDefaultGitHub_PullRequestForBranch(t *testing.T) {
	Getter = func() HttpGetter {
		return &FakeGetter{
			t: t,
			requests: map[string]string{
				"https://host/api/v3/repos/org/repo/pulls?head=org:no-pull-requests": "[]",
				"https://host/api/v3/repos/org/repo/pulls?head=org:pull-requests":    "[{}]",
			},
		}
	}

	gh := DefaultGitHub{}
	b, err := gh.PullRequestForBranch("host", "org", "repo", "no-pull-requests")
	assert.NoError(t, err)
	assert.False(t, b)

	b, err = gh.PullRequestForBranch("host", "org", "repo", "pull-requests")
	assert.NoError(t, err)
	assert.True(t, b)
}

type FakeGetter struct {
	t        *testing.T
	requests map[string]string
}

func (f *FakeGetter) Get(url string) (*http.Response, error) {
	s, ok := f.requests[url]
	if ok {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(s)),
		}, nil
	}
	return &http.Response{
		StatusCode: 404,
		Body:       ioutil.NopCloser(strings.NewReader("404")),
	}, nil
}

package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestDefaultGitHub_PullRequestForBranch(t *testing.T) {
	Getter = func() HttpGetter {
		return &FakeGetter{}
	}

	gh := DefaultGitHub{}
	b, err := gh.PullRequestForBranch("host", "org", "repo", "branch")
	assert.NoError(t, err)
	assert.False(t, b)
}

type FakeGetter struct {
}

func (f *FakeGetter) Get(url string) (*http.Response, error) {
	fmt.Println(url)
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("[]")),
	}, nil
}

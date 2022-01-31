package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultGitter_GetOrigin(t *testing.T) {
	Runner = func(command string, args ...string) (string, error) {
		assert.Equal(t, "git", command)
		assert.Equal(t, []string{"remote", "-v"}, args)
		return `origin	https://github.com/garethjevans/create-pr (fetch)
origin	https://github.com/garethjevans/create-pr (push)
fork	https://github.com/garethjevans/other (fetch)
fork	https://github.com/garethjevans/other (push)`, nil
	}
	g := DefaultGitter{}
	origin, err := g.GetOrigin()
	assert.NoError(t, err)
	assert.Equal(t, "https://github.com/garethjevans/create-pr", origin)
}

func TestDefaultGitter_HasLocalChanges(t *testing.T) {
	Runner = func(command string, args ...string) (string, error) {
		assert.Equal(t, "git", command)
		assert.Equal(t, []string{"status", "--porcelain"}, args)
		return `M somefile`, nil
	}
	g := DefaultGitter{}
	changes, err := g.HasLocalChanges()
	assert.NoError(t, err)
	assert.Equal(t, true, changes)
}

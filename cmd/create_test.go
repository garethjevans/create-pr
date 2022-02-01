package cmd

import (
	"fmt"
	"github.com/garethjevans/create-pr/pkg/pkgfakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePr_Run_NoLocalChanges(t *testing.T) {
	gitter := &pkgfakes.FakeGitter{}
	//github := &pkgfakes.FakeGitHub{}

	create := CreatePr{
		Gitter: gitter,
		//GitHub: github,
	}

	gitter.GetOriginReturns("https://github.com/myorg/myrepo", nil)
	gitter.HasLocalChangesReturns(false, nil)

	fmt.Println(create.Run())

	assert.Equal(t, 1, gitter.GetOriginCallCount())
	assert.Equal(t, 1, gitter.HasLocalChangesCallCount())
}

func TestCreatePr_Run_LocalChanges(t *testing.T) {
	gitter := &pkgfakes.FakeGitter{}
	github := &pkgfakes.FakeGitHub{}

	create := CreatePr{
		Gitter: gitter,
		GitHub: github,
	}

	gitter.GetOriginReturns("https://github.com/myorg/myrepo", nil)
	gitter.HasLocalChangesReturns(true, nil)

	github.DefaultBranchReturns("main", nil)
	github.PullRequestForBranchReturns(false, nil)

	fmt.Println(create.Run())

	assert.Equal(t, 1, gitter.GetOriginCallCount())
	assert.Equal(t, 1, gitter.HasLocalChangesCallCount())

	assert.Equal(t, 1, github.DefaultBranchCallCount())
	assert.Equal(t, 1, github.PullRequestForBranchCallCount())
}

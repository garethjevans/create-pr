/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/garethjevans/create-pr/pkg"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

type CreatePr struct {
	Command     *cobra.Command
	GithubToken string
	Branch      string
	Message     string
	Gitter      pkg.Gitter
	GitHub      pkg.GitHub
}

func NewCreatePr() CreatePr {
	pr := CreatePr{}
	pr.Command = &cobra.Command{
		Use:   "create",
		Short: "creates a pull request with the local changes",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			err := pr.Run()
			if err != nil {
				panic(err)
			}
		},
	}

	pr.Command.Flags().StringVarP(&pr.GithubToken, "github-token", "t", "", "The token to use when calling the GitHub API & pushing to git")
	pr.Command.Flags().StringVarP(&pr.Branch, "branch", "b", "", "The branch to use")
	pr.Command.Flags().StringVarP(&pr.Message, "message", "m", "", "The message to use when pushing the commit")

	return pr
}

func (c *CreatePr) Run() error {
	// TODO if the token hasn't been specified, we should check the environment variable
	// fall back on the gh config? ~/.config/gh/hosts.yml - can we find an entry for this host?

	fmt.Println("create called")
	// can we get information about this repo
	if c.Gitter == nil {
		c.Gitter = pkg.DefaultGitter{}
	}

	if c.GitHub == nil {
		c.GitHub = pkg.DefaultGitHub{}
	}

	remote, err := c.Gitter.GetOrigin()
	if err != nil {
		return err
	}
	fmt.Println(remote)
	remoteUrl, err := url.Parse(remote)
	if err != nil {
		return err
	}

	host := remoteUrl.Host
	parts := strings.Split(remoteUrl.Path, "/")
	org := parts[1]
	repo := parts[2]

	// determine if there are local changes
	changes, err := c.Gitter.HasLocalChanges()
	if err != nil {
		return err
	}

	if !changes {
		fmt.Println("There is nothing to do.")
		return nil
	}

	// what is the main branch for this repository
	defaultBranch, err := c.GitHub.DefaultBranch(host, org, repo)
	if err != nil {
		return err
	}

	fmt.Println(defaultBranch)

	// are there any pull requests for this branch
	pullRequestExists, err := c.GitHub.PullRequestForBranch(host, org, repo, c.Branch)
	if err != nil {
		return err
	}

	if pullRequestExists {
		fmt.Println("There is already a pull request for this branch")
	} else {
		fmt.Println("No pull requests exist for this branch")
	}

	// TODO what is the min to get this to work?
	// git config??
	// git checkout -b $branch
	// git add -A
	// git commit -m ""
	// git push origin $branch
	// create the pull request

	return nil
}

func init() {
	rootCmd.AddCommand(NewCreatePr().Command)
}

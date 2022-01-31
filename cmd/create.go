/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

type CreatePr struct {
	Command     *cobra.Command
	GithubToken string
	Branch      string
	Message     string
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
	pr.Command.Flags().StringVarP(&pr.GithubToken, "branch", "b", "", "The branch to use")
	pr.Command.Flags().StringVarP(&pr.GithubToken, "message", "m", "", "The message to use when pushing the commit")

	return pr
}

func (c *CreatePr) Run() error {
	fmt.Println("create called")
	// can we get information about this repo
	gitter := DefaultGitter{}
	remote, err := gitter.GetOrigin()
	if err != nil {
		return err
	}
	fmt.Println(remote)
	remoteUrl, err := url.Parse(remote)
	if err != nil {
		return err
	}

	host := remoteUrl.Host
	fmt.Println("host", host)
	parts := strings.Split(remoteUrl.Path, "/")
	org := parts[1]
	repo := parts[2]

	fmt.Println("org", org)
	fmt.Println("repo", repo)

	// determine if there are local changes
	changes, err := gitter.HasLocalChanges()
	if err != nil {
		return err
	}
	fmt.Println(changes)

	// what is the main branch for this repository

	// are there any pull requests for this branch
	gh := DefaultGitHub{}
	b, err := gh.PullRequestForBranch(host, org, repo, "branch")
	if err != nil {
		return err
	}
	fmt.Println(b)
	return nil
}

func init() {
	rootCmd.AddCommand(NewCreatePr().Command)
}

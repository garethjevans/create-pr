/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
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
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	fmt.Println(host)

	// determine if there are local changes
	changes, err := gitter.HasLocalChanges()
	if err != nil {
		return err
	}
	fmt.Println(changes)

	// what is the main branch for this repository

	// are there any pull requests for this branch
	gh := DefaultGitHub{}
	b, err := gh.PullRequestForBranch(host, "garethjevans", "create-pr", "branch")
	if err != nil {
		return err
	}
	fmt.Println(b)
	return nil
}

func init() {
	rootCmd.AddCommand(NewCreatePr().Command)
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	return nil
}

func init() {
	rootCmd.AddCommand(NewCreatePr().Command)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//
}

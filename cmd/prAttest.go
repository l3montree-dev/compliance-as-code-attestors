/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/l3montree/compliance-as-code-attestors/cross_pr_request"
	"github.com/spf13/cobra"
)

// prAttestCmd represents the prAttest command
var prAttestCmd = &cobra.Command{
	Use:   "prAttest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		initRepoTitle, err := cmd.Flags().GetString("initRepoTitle")
		initRepoNumber, err := cmd.Flags().GetInt("initRepoNumber")
		rawRepos, err := cmd.Flags().GetStringArray("repos")

		if err != nil {
			fmt.Print("Error")
			return
		}

		var repos []string
		for _, r := range rawRepos {
			for _, repo := range strings.Split(r, ",") {
				repo = strings.TrimSpace(repo)
				if repo == "" {
					continue
				}
				repos = append(repos, repo)
			}
		}

		cross_pr_request.CrossPRrequest(repos, initRepoNumber, initRepoTitle)
	},
}

func init() {
	rootCmd.AddCommand(prAttestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	prAttestCmd.Flags().Int("initRepoNumber", 3, "insert pullRequest Number here")
	// prAttestCmd.Flags().StringArrayVar(["test"], "test", "test2")
	prAttestCmd.Flags().StringArray("repos", []string{"l3montree-dev/devguard,l3montree-dev/devguard-documentation"}, "Comma-separated list of repositories or repeat --repos for each.")
	prAttestCmd.Flags().String("initRepoTitle", "1277 organization wide dependency search", "Name of the Pull Request")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prAttestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

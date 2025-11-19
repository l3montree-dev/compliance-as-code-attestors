/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		repos, err := cmd.Flags().GetStringArray("repos")

		if err != nil {
			fmt.Print("Error")
			return
		}
		fmt.Printf("%s", initRepoTitle)
		fmt.Printf("%d", initRepoNumber)
		fmt.Printf(repos[0])
	},
}

func init() {
	rootCmd.AddCommand(prAttestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	prAttestCmd.Flags().Int("initRepoNumber", 0, "insert pullRequest Number here")
	// prAttestCmd.Flags().StringArrayVar(["test"], "test", "test2")
	prAttestCmd.Flags().StringArray("repos", []string{"repository"}, `{"exampleRepo1","exampleRepo2"}`)
	prAttestCmd.Flags().String("initRepoTitle", "Example Title", "Name of the Pull Request")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prAttestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

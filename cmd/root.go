package cmd

import (
	"github.com/ihribernik/aoc-cli/internal/container"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "aoc-cli",
	Short:         "Advent of Code CLI runner",
	Long:          `Advent of Code CLI is a command-line tool to run Advent of Code solutions.`,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	configuredContainer, err := container.New()
	if err != nil {
		return mapStartupError(err)
	}

	appContainer = configuredContainer
	return rootCmd.Execute()
}

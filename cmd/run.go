package cmd

import (
	"fmt"

	"github.com/ihribernik/aoc-cli/internal/container"
	runusecase "github.com/ihribernik/aoc-cli/internal/run"
	"github.com/spf13/cobra"
)

var appContainer container.Container

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:          "run --year <year> --day <day>",
	Short:        "Execute the solution for a specific Advent of Code puzzle",
	Args:         cobra.NoArgs,
	RunE:         runE,
	SilenceUsage: true,
}

func runE(cmd *cobra.Command, args []string) error {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		return fmt.Errorf("invalid --year value: %w", err)
	}

	day, err := cmd.Flags().GetInt("day")
	if err != nil {
		return fmt.Errorf("invalid --day value: %w", err)
	}

	if day < 1 || day > 25 {
		return fmt.Errorf("invalid --day %d: expected a value between 1 and 25", day)
	}

	if appContainer == nil || appContainer.GetRunner() == nil {
		return newCLIError("application runner is not configured", runusecase.ErrRunnerNotConfigured)
	}

	result, err := appContainer.GetRunner().Execute(year, day)

	if err != nil {
		return mapRunError(year, day, err)
	}

	fmt.Println("Solution part 1:", result.Part1)
	fmt.Println("Solution part 2:", result.Part2)
	return nil
}

func init() {
	runCmd.Flags().IntP("year", "y", 0, "Year to execute (e.g. 2015)")
	runCmd.Flags().IntP("day", "d", 0, "Day to execute (e.g. 6)")
	_ = runCmd.MarkFlagRequired("year")
	_ = runCmd.MarkFlagRequired("day")
	rootCmd.AddCommand(runCmd)
}

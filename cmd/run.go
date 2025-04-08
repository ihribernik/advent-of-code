/*
Copyright © 2025 ivan hribernik cihribernik@gmail.com
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ihribernik/advent-of-code/internal/solutions/y2015"
	"github.com/ihribernik/advent-of-code/pkg/common/directions"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [year] [day]",
	Short: "Ejecuta el ejercicio para el Año dia indicado",
	Args:  cobra.ExactArgs(2),
	Run:   runner,
}

func runner(cmd *cobra.Command, args []string) {
	year, err := strconv.Atoi(args[0])
	if err != nil {
		panic("Error: year must be a string")
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		panic("Error: day must be a string")
	}

	directionA := directions.Direction{1, 1}
	directionB := directions.Direction{1, -1}

	result := directionA > directionB
	panic("a is greather than b %v", result)

	result, err := runSolution(year, day)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The result is %v", result)
}

func runSolution(year int, day int) (int, error) {
	switch year {
	case 2015:
		executedYear, err := runYear2015(day)
		return executedYear, err
	default:
		return 0, fmt.Errorf("error: no se encontro el anio")
	}
}

func runYear2015(day int) (int, error) {
	switch day {
	case 1:
		executedDay, err := y2015.Day01()
		return executedDay, err
	default:
		return 0, fmt.Errorf("no se Encontro el dia ejecutar")
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

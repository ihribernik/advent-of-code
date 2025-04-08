/*
Copyright © 2025 ivan hribernik cihribernik@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ihribernik/advent-of-code/internal/solutions"
	_ "github.com/ihribernik/advent-of-code/internal/solutions/y2015"
	"github.com/ihribernik/advent-of-code/pkg/common"
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
		log.Fatal("Error: year must be a string")
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("Error: day must be a string")
	}

	solver, ok := solutions.GetSolver(year, day)
	if !ok {
		log.Fatalf("No se encontró solución para el año %d, día %d", year, day)
	}

	runner := common.Runner{Year: year, Day: day}
	input, err := runner.Parse()

	result1, err := solver.SolvePart1(input)
	result2, err := solver.SolvePart2(input)
	fmt.Println("Parte 1:", result1)
	fmt.Println("Parte 2:", result2)
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

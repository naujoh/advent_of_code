// Package internal
package internal

import (
	"flag"
	"log/slog"

	"github.com/naujoh/aoc/internal/solutions"
)

type terminalParams struct {
	verbose      bool
	useTestInput bool
	puzzleID     string
}

func Solve() {
	params := readParams()
	if params.verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Info("Loading", "puzzle", params.puzzleID)
	solution := solutions.LoadSolution(
		params.puzzleID,
		params.useTestInput,
	)
	puzzle := solution.GetPuzzle()

	slog.Info("Running puzzle", "name", puzzle.Name)

	slog.Info("Running first part solution...")
	firstPartResult := handleResult(solution.SolveFirstPart())
	slog.Info("====")

	slog.Info("Running second part solution...")
	secondPartResult := handleResult(solution.SolveSecondPart())
	slog.Info("====")

	slog.Info("First part:", "solution", firstPartResult)
	slog.Info("Second part:", "solution", secondPartResult)

	slog.Info("Happy Advent! Hope you get all the gold stars *")
}

func handleResult(result string) string {
	if result == "" {
		return "<unsolved>"
	}
	return result
}

func readParams() terminalParams {
	params := terminalParams{verbose: false}

	flag.BoolVar(&params.verbose, "v", false, "Enable verbose output (short)")
	flag.BoolVar(&params.verbose, "verbose", false, "Enable verbose output (short)")
	flag.BoolVar(
		&params.useTestInput,
		"t",
		false,
		"Enable to use test file, uses puzzle file if not specified",
	)
	flag.StringVar(&params.puzzleID, "puzzle", "", "The puzzle identifier to run its solution")
	flag.Parse()

	return params
}

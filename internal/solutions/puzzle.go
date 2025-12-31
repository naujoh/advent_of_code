// Package solutions
package solutions

import (
	"embed"
	"log/slog"
	"os"

	"github.com/naujoh/aoc/pkg/utils"
)

type PuzzleSolver interface {
	GetPuzzle() BasePuzzle
	SolveFirstPart() string
	SolveSecondPart() string
}

type BasePuzzle struct {
	PuzzleID     string
	Name         string
	UseTestInput bool
	PuzzleInput  string
	TestInput    string
}

type PuzzleFactoryFn func(useTestInput bool) PuzzleSolver

var PuzzleRegistry = make(map[string]PuzzleFactoryFn)

func AddPuzzle(puzzleID string, factoryFn PuzzleFactoryFn) {
	if _, ok := PuzzleRegistry[puzzleID]; ok {
		panic("Puzzle already loaded")
	}
	PuzzleRegistry[puzzleID] = factoryFn
}

func LoadSolution(puzzleID string, useTestInput bool) PuzzleSolver {
	if _, ok := PuzzleRegistry[puzzleID]; !ok {
		slog.Error("Puzzle not found, I feel sorry for you :(")
		os.Exit(1)
	}
	return PuzzleRegistry[puzzleID](useTestInput)
}

func (bp *BasePuzzle) LoadPuzzleInput(fs embed.FS) string {
	slog.Debug("Is using test input?", "useTestInput", bp.UseTestInput)
	if bp.UseTestInput {
		return utils.ReadInputFile(fs, bp.TestInput)
	} else {
		return utils.ReadInputFile(fs, bp.PuzzleInput)
	}
}

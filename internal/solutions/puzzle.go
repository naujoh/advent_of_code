// Package solutions
package solutions

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/naujoh/aoc/pkg/utils"
)

const (
	puzzleInputURLBase      = "https://adventofcode.com/%d/day/%d/input"
	puzzleDownloaderTimeout = 10 * time.Second
)

type PuzzleSolver interface {
	GetPuzzle() BasePuzzle
	SolveFirstPart() string
	SolveSecondPart() string
}

type BasePuzzle struct {
	PuzzleID     string
	Name         string
	Year         int
	Day          int
	UseTestInput bool
	PuzzleInput  string
	TestInput    string
}

type PuzzleFactoryFn func(useTestInput bool) PuzzleSolver

var PuzzleRegistry = make(map[string]PuzzleFactoryFn)

func Create(name string, year string, day string) BasePuzzle {
	return BasePuzzle{
		PuzzleID:    year + "-" + day,
		Name:        name,
		Year:        utils.StrToInt(year),
		Day:         utils.StrToInt(day),
		PuzzleInput: "day" + day + ".puzzle",
		TestInput:   "day" + day + ".test",
	}
}

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

func (bp *BasePuzzle) LoadPuzzleInput() string {
	slog.Debug("Is using test input?", "useTestInput", bp.UseTestInput)
	puzzleFilePath := bp.GetPuzzleFilePath()

	if !bp.UseTestInput && !utils.FileExists(puzzleFilePath) {
		slog.Warn("Puzzle input not found")
		bp.DownloadPuzzleInputFile()
	}
	return utils.ReadInputFile(puzzleFilePath)
}

func (bp *BasePuzzle) DownloadPuzzleInputFile() {
	slog.Info("Trying to download puzzle input for puzzle", "puzzle", bp)
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		slog.Error("Cannot find cookie on AOC_SESSION environment variable")
		os.Exit(1)
	}

	puzzleInputURL := fmt.Sprintf(puzzleInputURLBase, bp.Year, bp.Day)

	client := &http.Client{
		Timeout: puzzleDownloaderTimeout,
	}

	req, err := http.NewRequest(http.MethodGet, puzzleInputURL, nil)
	if err != nil {

		slog.Error("Cannot download puzzle input, error creating request", "error", err)
		os.Exit(1)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, err := client.Do(req)
	if err != nil {
		slog.Error("Cannot download puzzle input, error getting response", "error", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		slog.Error("Cannot download puzzle input, request failed", "reqStatus", res.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Cannot download puzzle input, error reading response body", "error", err)
		os.Exit(1)
	}

	puzzleInputFilePath := bp.GetPuzzleFilePath()
	if os.WriteFile(puzzleInputFilePath, body, 0644); err != nil {
		slog.Error("Cannot save puzzle input, error", "error", err)
		os.Exit(1)
	}

	slog.Info("Puzzle input downloaded at", "path", puzzleInputFilePath)
}

func (bp *BasePuzzle) GetPuzzleFilePath() string {
	wd, err := os.Getwd()
	if err != nil {
		slog.Error(
			"Cannot get puzzle file path, cannot determine working directory",
			"error", err,
		)
		os.Exit(1)
	}

	fileName := bp.PuzzleInput
	if bp.UseTestInput {
		fileName = bp.TestInput
	}

	puzzleInputFilePath := filepath.Join(
		wd,
		"internal",
		"solutions",
		fmt.Sprintf("year%d", bp.Year),
		"inputs",
		fileName,
	)

	return puzzleInputFilePath
}

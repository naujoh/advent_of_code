package year2025

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/naujoh/aoc/internal/solutions"
)

type Day04 struct{ solutions.BasePuzzle }

func init() {
	name := "Printing department"
	stringDay := "04"
	puzzle := &Day04{BasePuzzle: solutions.Create(name, year, stringDay)}
	solutions.AddPuzzle(
		puzzle.PuzzleID, func(useTestInput bool) solutions.PuzzleSolver {
			puzzle.UseTestInput = useTestInput
			return puzzle
		},
	)
}

func (p *Day04) GetPuzzle() solutions.BasePuzzle { return p.BasePuzzle }

func (p *Day04) SolveFirstPart() string {
	grid := p.getRollPaperGrid()
	accessibleRollsCount := 0

	for i, row := range grid {
		slog.Debug("Row", "idx", i)
		for j := 0; j < len(row); j++ {
			if string(grid[i][j]) != "@" {
				continue
			}
			slog.Debug(
				"Checking adjacents for slot at",
				"row", i,
				"column", j,
				"val", string(grid[i][j]),
			)
			adjacentRollSum := 0
			directions := [][]int{
				{i, j - 1},     // left
				{i, j + 1},     // right
				{i - 1, j - 1}, // top left
				{i - 1, j},     // top
				{i - 1, j + 1}, // top right
				{i + 1, j - 1}, // bottom left
				{i + 1, j},     // bottom
				{i + 1, j + 1}, // bottom right
			}
			for _, p := range directions {
				x, y := p[0], p[1]

				// Check in boundaries
				if x < 0 || x > len(grid)-1 || y < 0 || y > len(row)-1 {
					continue
				}
				slot := string(grid[x][y])
				if slot == "@" {
					slog.Debug("Adjacent paper roll found at", "row", x, "column", y)
					adjacentRollSum++
				}
			}

			if adjacentRollSum < 4 {
				accessibleRollsCount++
				slog.Debug(
					"Adjacent rolls count < 4 found",
					"total", accessibleRollsCount,
				)
			}
		}
	}
	return strconv.Itoa(accessibleRollsCount)
}

func (p *Day04) SolveSecondPart() string {
	return ""
}

// Solution implementation

func (p *Day04) getRollPaperGrid() []string {
	input := p.LoadPuzzleInput()
	input = strings.TrimSpace(input)
	grid := []string{}

	grid = append(grid, strings.Split(input, "\n")...)

	return grid
}

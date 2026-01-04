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

	p.printGrid(grid)
	removed := strconv.Itoa(p.removePaperRollsFromGrid(grid))
	p.printGrid(grid)

	return removed
}

func (p *Day04) SolveSecondPart() string {
	grid := p.getRollPaperGrid()
	var removed, totalRemoved int

	for {
		slog.Debug("Grid before remove")
		p.printGrid(grid)
		removed = p.removePaperRollsFromGrid(grid)
		slog.Debug("Grid after remove")
		p.printGrid(grid)
		totalRemoved += removed

		if removed <= 0 {
			break
		}
	}

	return strconv.Itoa(totalRemoved)
}

// Solution implementation

func (p *Day04) removePaperRollsFromGrid(grid [][]string) int {
	rollsToRemove := [][]int{}

	for i, row := range grid {
		slog.Debug("Row", "idx", i)
		for j := range row {
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
				rollsToRemove = append(rollsToRemove, []int{i, j})
				slog.Debug(
					"Adjacent rolls count < 4 found",
					"total", len(rollsToRemove),
				)
			}
		}
	}

	for _, p := range rollsToRemove {
		grid[p[0]][p[1]] = "x"
	}

	return len(rollsToRemove)
}

func (p *Day04) printGrid(grid [][]string) {
	for i, row := range grid {
		slog.Debug("", "i", i, "row", row)
	}
}

func (p *Day04) getRollPaperGrid() [][]string {
	input := p.LoadPuzzleInput()
	input = strings.TrimSpace(input)
	grid := [][]string{}

	for i, row := range strings.Split(input, "\n") {
		grid = append(grid, []string{})
		for _, val := range row {
			grid[i] = append(grid[i], string(val))
		}
	}

	return grid
}

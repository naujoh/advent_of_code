// Package year2025
package year2025

import (
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/naujoh/aoc/internal/solutions"
	"github.com/naujoh/aoc/pkg/utils"
)

type Day01 struct{ solutions.BasePuzzle }

func init() {
	name := "Secret Entrance"
	stringDay := "01"

	puzzle := &Day01{
		BasePuzzle: solutions.BasePuzzle{
			Name:        name,
			PuzzleID:    year + "-" + stringDay,
			PuzzleInput: "inputs/day" + stringDay + ".puzzle",
			TestInput:   "inputs/day" + stringDay + ".test",
		},
	}
	solutions.AddPuzzle(
		puzzle.PuzzleID,
		func(useTestInput bool) solutions.PuzzleSolver {
			puzzle.UseTestInput = useTestInput
			return puzzle
		},
	)
}

func (p *Day01) GetPuzzle() solutions.BasePuzzle { return p.BasePuzzle }

func (p *Day01) SolveFirstPart() string {
	return p.solve(false)
}

func (p *Day01) SolveSecondPart() string {
	return p.solve(true)
}

// Solution implementation
type rotation struct {
	direction rune
	clicks    int
}
type dial struct {
	length      int
	position    int
	combination int
}

func (p *Day01) solve(isSecondPart bool) string {
	rotations := p.getRotationList()
	doorDial := dial{100, 50, 0}

	for i, r := range rotations {
		clicks := r.clicks

		if r.clicks >= doorDial.length {
			if isSecondPart {
				doorDial.combination += r.clicks / doorDial.length
			}
			clicks %= doorDial.length
		}

		switch r.direction {
		case 'L':
			doorDial = rotateDialToLeft(doorDial, clicks, isSecondPart)
		case 'R':
			doorDial = rotateDialToRight(doorDial, clicks, isSecondPart)
		default:
			slog.Error("No valid direction found", "idx", i, "rotation", r)
			os.Exit(1)
		}

		if doorDial.position == 0 {
			doorDial.combination++
		}

		slog.Debug(
			"End rotation",
			"rotationNumber", i+1,
			"direction", string(r.direction),
			"clicks", r.clicks,
			"dialPos", doorDial.position,
			"combination", doorDial.combination,
		)
	}

	return strconv.Itoa(doorDial.combination)
}

func rotateDialToLeft(doorDial dial, clicks int, isSecondPart bool) dial {
	doorDial.position -= clicks

	if isSecondPart &&
		doorDial.position < 0 &&
		doorDial.position+clicks > 0 {

		doorDial.combination++
	}

	if doorDial.position < 0 {
		doorDial.position = doorDial.length - (-doorDial.position)
	}
	return doorDial
}

func rotateDialToRight(doorDial dial, clicks int, isSecondPart bool) dial {
	doorDial.position += clicks

	if isSecondPart &&
		doorDial.position > doorDial.length &&
		doorDial.position-clicks < doorDial.length {

		doorDial.combination++
	}

	if doorDial.position >= doorDial.length {
		doorDial.position = doorDial.position - doorDial.length
	}
	return doorDial
}

func (p *Day01) getRotationList() []rotation {
	input := p.LoadPuzzleInput(inputsFS)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	rotations := []rotation{}

	for i := range len(lines) {
		rotations = append(
			rotations,
			rotation{
				direction: rune(lines[i][0]),
				clicks:    utils.StrToInt(lines[i][1:]),
			},
		)
	}

	return rotations
}

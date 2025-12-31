package year2025

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/naujoh/aoc/internal/solutions"
	"github.com/naujoh/aoc/pkg/utils"
)

type Day03 struct{ solutions.BasePuzzle }

func init() {
	name := "Lobby"
	stringDay := "03"

	puzzle := &Day03{
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

func (p *Day03) GetPuzzle() solutions.BasePuzzle { return p.BasePuzzle }

func (p *Day03) SolveFirstPart() string {
	batteryBanks := p.getBaterryBanks()
	joltageSum := 0

	for _, bank := range batteryBanks {
		maxIds := []int{0}
		slog.Debug("Bank", "value", bank)
		for i := 1; i < len(bank); i++ {
			if bank[maxIds[len(maxIds)-1]] <= bank[i] {
				maxIds = append(maxIds, i)
			}
		}
		slog.Debug("Max idxs", "value", maxIds)

		if len(maxIds) > 1 &&
			(maxIds[len(maxIds)-1] == len(bank)-1 ||
				bank[maxIds[len(maxIds)-1]] == bank[maxIds[len(maxIds)-2]]) {
			slog.Debug(
				"Greater pair",
				"l",
				bank[maxIds[len(maxIds)-2]],
				"r",
				bank[maxIds[len(maxIds)-1]],
			)
			number := strconv.Itoa(
				bank[maxIds[len(maxIds)-2]],
			) + strconv.Itoa(
				bank[maxIds[len(maxIds)-1]],
			)
			joltageSum += utils.StrToInt(number)

		} else {
			maxRight := 0
			for i := maxIds[len(maxIds)-1] + 1; i < len(bank); i++ {
				maxRight = max(maxRight, bank[i])
			}
			slog.Debug(
				"Greater pair",
				"l",
				bank[maxIds[len(maxIds)-1]],
				"r",
				maxRight,
			)
			number := strconv.Itoa(bank[maxIds[len(maxIds)-1]]) + strconv.Itoa(maxRight)
			joltageSum += utils.StrToInt(number)
		}

	}
	return strconv.Itoa(joltageSum)
}

func (p *Day03) SolveSecondPart() string {
	return ""
}

// Solution implementation

func (p *Day03) getBaterryBanks() [][]int {
	input := p.LoadPuzzleInput(inputsFS)
	input = strings.TrimSpace(input)
	var banks [][]int

	for i, b := range strings.Split(input, "\n") {
		banks = append(banks, []int{})
		for j := 0; j < len(b); j++ {
			banks[i] = append(banks[i], utils.StrToInt(string(b[j])))
		}
	}

	return banks
}

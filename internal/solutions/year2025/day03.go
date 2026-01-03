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
	puzzle := &Day03{BasePuzzle: solutions.Create(name, year, stringDay)}
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
	return p.solve(2)
}

func (p *Day03) SolveSecondPart() string {
	return p.solve(12)
}

// Solution implementation

func (p *Day03) solve(k int) string {
	batteryBanks := p.getBaterryBanks()
	totalJoltage := 0
	for i, bank := range batteryBanks {
		l := len(bank)
		number := []int{}
		outputNumber := []string{}
		digitPosition := 0

		slog.Debug("Bank ", "No", i, "bank", bank)
		for j, batteryValue := range bank {
			for len(number) > 0 && number[len(number)-1] < batteryValue {
				if l-j <= k { // i is at the las K numbers of bank
					relativePosition := k - (l - j) + 1
					if relativePosition > digitPosition {
						break
					}
				}

				number = number[:len(number)-1]
				digitPosition--
			}

			if len(number) < k {
				number = append(number, batteryValue)
				digitPosition++
				slog.Debug("Number", "n", number)
			}
		}
		slog.Debug("----")

		for _, d := range number {
			outputNumber = append(outputNumber, strconv.Itoa(d))
		}
		totalJoltage += utils.StrToInt(strings.Join(outputNumber, ""))
	}
	return strconv.Itoa(totalJoltage)
}

func (p *Day03) getBaterryBanks() [][]int {
	input := p.LoadPuzzleInput()
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

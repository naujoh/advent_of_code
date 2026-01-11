package year2025

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/naujoh/aoc/internal/solutions"
	"github.com/naujoh/aoc/pkg/utils"
)

type Day05 struct{ solutions.BasePuzzle }

func init() {
	name := "Cafeteria"
	stringDay := "05"
	puzzle := &Day05{BasePuzzle: solutions.Create(name, year, stringDay)}
	solutions.AddPuzzle(
		puzzle.PuzzleID, func(useTestInput bool) solutions.PuzzleSolver {
			puzzle.UseTestInput = useTestInput
			return puzzle
		},
	)
}

func (p *Day05) GetPuzzle() solutions.BasePuzzle { return p.BasePuzzle }

func (p *Day05) SolveFirstPart() string {
	ingredientRanges, ingredients := p.readIngredientsDB()
	freshIngredientsCount := 0

	for _, ingredientRange := range ingredientRanges {
		l, r := ingredientRange[0], ingredientRange[1]
		slog.Debug("Ingredient range", "l", l, "r", r)

		remainingIngredients := []int{}

		for _, i := range ingredients {
			slog.Debug("Ingredient ID", "id", i)
			if i >= l && i <= r {
				freshIngredientsCount++
				continue
			}
			remainingIngredients = append(remainingIngredients, i)
		}

		ingredients = remainingIngredients
	}
	return strconv.Itoa(freshIngredientsCount)
}

func (p *Day05) SolveSecondPart() string {
	return ""
}

// Solution implementation

func (p *Day05) readIngredientsDB() ([][]int, []int) {
	input := p.LoadPuzzleInput()
	input = strings.TrimSpace(input)

	ranges := [][]int{}
	ingredients := []int{}

	storingRanges := true

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			storingRanges = false
			continue
		}
		if !storingRanges {
			// Store ingredient ids
			ingredients = append(ingredients, utils.StrToInt(line))
			continue
		}
		// Store ingredient ranges
		ingredientRange := strings.Split(line, "-")
		ranges = append(ranges, []int{
			utils.StrToInt(ingredientRange[0]),
			utils.StrToInt(ingredientRange[1]),
		})
	}

	return ranges, ingredients
}

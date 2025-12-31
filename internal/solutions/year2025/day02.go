package year2025

import (
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/naujoh/aoc/internal/solutions"
	"github.com/naujoh/aoc/pkg/utils"
)

type Day02 struct{ solutions.BasePuzzle }

func init() {
	name := "Gift Shop"
	stringDay := "02"

	puzzle := &Day02{BasePuzzle: solutions.Create(name, year, stringDay)}
	solutions.AddPuzzle(
		puzzle.PuzzleID,
		func(useTestInput bool) solutions.PuzzleSolver {
			puzzle.UseTestInput = useTestInput
			return puzzle
		},
	)
}

func (p *Day02) GetPuzzle() solutions.BasePuzzle { return p.BasePuzzle }

func (p *Day02) SolveFirstPart() string {
	var invalidIDsSum int64

	for _, idRange := range p.getIDRanges() {
		slog.Debug("Range", "left", idRange.left, "right", idRange.right)
		seenIDs := map[int64]struct{}{}
		groupDef := groupsDefinition{groupsCount: 2}
		limitID := idRange.right

		if idRange.leftDigitsCount%2 == 0 {
			groupDef.groupSize = idRange.leftDigitsCount / 2

			if idRange.leftDigitsCount != idRange.rightDigitsCount {
				limitID = int64(math.Pow10(idRange.leftDigitsCount))
			}

			invalidIDsSum += sumInvalidIDs(idRange.left, limitID, groupDef, seenIDs)
		}

		if idRange.leftDigitsCount != idRange.rightDigitsCount &&
			idRange.rightDigitsCount%2 == 0 {

			// Get sum of invalid ids from the right range value
			groupDef.groupSize = idRange.rightDigitsCount / 2
			invalidIDsSum += sumInvalidIDs(
				int64(math.Pow10(idRange.rightDigitsCount-1)),
				idRange.right,
				groupDef,
				seenIDs,
			)
		}

		slog.Debug("---")
	}
	return strconv.Itoa(int(invalidIDsSum))
}

func (p *Day02) SolveSecondPart() string {
	var invalidIDsSum int64

	for _, idRange := range p.getIDRanges() {
		slog.Debug("Range", "left", idRange.left, "right", idRange.right)
		seenIDs := map[int64]struct{}{}
		groupDef := groupsDefinition{groupSize: 1}
		limitID := idRange.right

		if idRange.leftDigitsCount != idRange.rightDigitsCount {
			limitID = int64(math.Pow10(idRange.leftDigitsCount))
		}

		groupDef.groupsCount = idRange.leftDigitsCount
		invalidIDsSum += sumInvalidIDs(idRange.left, limitID, groupDef, seenIDs)

		if idRange.leftDigitsCount != idRange.rightDigitsCount {
			// Get sum of invalid ids from the right range value
			groupDef.groupsCount = idRange.rightDigitsCount
			invalidIDsSum += sumInvalidIDs(
				int64(math.Pow10(idRange.rightDigitsCount-1)),
				idRange.right,
				groupDef,
				seenIDs,
			)
		}
		slog.Debug("---")
	}
	return strconv.Itoa(int(invalidIDsSum))
}

// Solution implementation
type idRange struct {
	left             int64
	right            int64
	leftDigitsCount  int
	rightDigitsCount int
}

type groupsDefinition struct {
	groupSize   int
	groupsCount int
}

func getNextInvalidID(x int64, k int, m int) (int64, int64) {
	// From formula to get the next repeated block of numbers with pattern ABAB, ABCABC
	// k = length of repeated block
	// m = times the block of numbers should be repeated
	// B = repeated block of numbers with pattern ABAB, ABCABC, ...
	// multiplier = ((10^km) - 1) / ((10^k) - 1)
	// N = B * multiplier
	// BNext = x / multiplier
	// NNext = BNext * multiplier

	multiplier := int64(math.Ceil((math.Pow10(int(k*m)) - 1) / (math.Pow10(int(k)) - 1)))
	slog.Debug("multiplier", "val", multiplier)
	BNext := int64(math.Ceil(float64(x) / float64(multiplier)))
	return BNext * int64(multiplier), int64(multiplier)
}

func sumInvalidIDs(
	startID int64, limitID int64, groupDef groupsDefinition, seenIDs map[int64]struct{},
) int64 {
	var invalidID int64
	var offset int64
	var invalidIDsSum int64

	startIDDigitsCount := len(strconv.Itoa(int(startID)))

	slog.Debug("startID", "val", startID)

	for groupDef.groupsCount >= 2 {
		if startIDDigitsCount%groupDef.groupSize == 0 {

			invalidID, offset = getNextInvalidID(
				startID, groupDef.groupSize, groupDef.groupsCount,
			)

			for invalidID <= limitID {
				if _, ok := seenIDs[invalidID]; !ok {
					invalidIDsSum += invalidID
					seenIDs[invalidID] = struct{}{}
				}

				slog.Debug("Generated invalid id", "invalidID", invalidID)
				invalidID += offset
			}
		}
		slog.Debug(
			"Group definition",
			"groupSize", groupDef.groupSize,
			"groupsAmount", groupDef.groupsCount,
		)
		groupDef.groupSize++
		groupDef.groupsCount = startIDDigitsCount / groupDef.groupSize
	}
	return invalidIDsSum
}

func (p *Day02) getIDRanges() []idRange {
	input := p.LoadPuzzleInput()
	input = strings.TrimSpace(input)
	ranges := []idRange{}

	for r := range strings.SplitSeq(input, ",") {
		rSplitted := strings.Split(r, "-")
		ranges = append(
			ranges,
			idRange{
				left:             int64(utils.StrToInt(rSplitted[0])),
				right:            int64(utils.StrToInt(rSplitted[1])),
				leftDigitsCount:  len(rSplitted[0]),
				rightDigitsCount: len(rSplitted[1]),
			},
		)
	}

	return ranges
}

// Package utils contains a collection of utilities functions
// that help to build puzzle solutions easily
package utils

import (
	"fmt"
	"strconv"
)

func StrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("error converting %s to int: %w", s, err))
	}
	return num
}

package utils

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func SortSlice[T constraints.Ordered](s []T) {
	slices.Sort(s)
}

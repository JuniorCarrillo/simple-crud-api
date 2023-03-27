package tools

import (
	"math"
)

func ToPages(count, limit int64) int64 {
	countFloat := float64(count)
	limitFloat := float64(limit)
	estimate := countFloat / limitFloat
	return int64(math.Ceil(estimate))
}

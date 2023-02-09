package gmath

import (
	"math"

	"github.com/ywh147906/load-test/common/utils/generic"
)

func CeilTo[T generic.Number](v float64) T {
	return T(math.Ceil(v))
}

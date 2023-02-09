package utils

import (
	"fmt"
	"math"

	"github.com/ywh147906/load-test/common/values"
)

func InRange(_range []values.Integer, val values.Integer) bool {
	if len(_range) < 2 {
		panic(fmt.Sprintf("range[%v] len < 2", _range))
	}
	if _range[1] == -1 {
		_range[1] = math.MaxInt64
	}
	return val >= _range[0] && val <= _range[1]
}

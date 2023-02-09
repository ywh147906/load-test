package enum

import "github.com/ywh147906/load-test/common/values"

type AttrShowType = values.Integer

const (
	Direct  AttrShowType = 1
	Percent AttrShowType = 2
)

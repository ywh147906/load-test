package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/ywh147906/load-test/common/values/env"
)

func GetRuleName() string {
	rule := strings.TrimSpace(os.Getenv(env.RULE_TAGE))
	if rule == "" {
		panic(fmt.Sprintf("无效的规则名"))
	}
	return rule
}

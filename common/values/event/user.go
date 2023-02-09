package event

import "github.com/ywh147906/load-test/common/values"

var UserLogin values.EventName = "UserLoginIn"

type UserLoginData struct {
	RoleId values.RoleId
}

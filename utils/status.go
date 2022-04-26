package CustomStatus

import "errors"

var UserNotFound error
var ExistUser error
var WrongPasswd error

func InitCustomStatus() {
	UserNotFound = errors.New("UserNotFound")
	ExistUser = errors.New("ExistUser")
	WrongPasswd = errors.New("WrongPasswd")
}

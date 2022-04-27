package CustomStatus

import "errors"

var UserNotFound error
var ExistUser error
var WrongPasswd error

var CustomerNotFound error
var ExistCustomer error

func InitCustomStatus() {
	UserNotFound = errors.New("UserNotFound")
	ExistUser = errors.New("ExistUser")
	WrongPasswd = errors.New("WrongPasswd")

	CustomerNotFound = errors.New("CustomerNotFound")
	ExistCustomer = errors.New("ExistCustomer")
}

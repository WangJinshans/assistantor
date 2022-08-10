package common

import "errors"

const (
	OrderPayed   = "payed"
	OrderCancel  = "canceled"
	OrderCreated = "created"
	OrderExpired = "expired"
)

const (
	RegularOrderType = iota
	VipOrderType
)

func IsPreLockNeeded(orderType int) (res bool, err error) {
	switch orderType {
	case RegularOrderType:
		res = true
		return
	case VipOrderType:
		return
	default:
		err = errors.New(OrderTypeError)
		return
	}
}

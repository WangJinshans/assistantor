package utils

import (
	"fmt"
	"time"
)

func GenerateOrderId() (orderId string) {
	orderId = fmt.Sprintf("%d-%s", time.Now().Unix(), GetUuidString())
	return
}

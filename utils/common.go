package utils

import (
	"fmt"
	"time"
)

func GenerateUniqueId() (orderId string) {
	orderId = fmt.Sprintf("%d-%s", time.Now().Unix(), GetUuidString())
	return
}

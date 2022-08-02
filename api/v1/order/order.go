package order

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateOrder(ctx *gin.Context) {
	js := make(map[string]interface{}) //注意该结构接受的内容
	ctx.BindJSON(&js)
	orderId := js["order_id"]
	log.Info().Msgf("order id is: %s", orderId)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

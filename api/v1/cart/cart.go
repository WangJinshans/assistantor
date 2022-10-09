package cart

import (
	"assistantor/common"
	"github.com/gin-gonic/gin"
)

// cart 默认存在 使用userId 只要添加绑定记录到本表, 即说明有商品在购物车中, 否则空空如也
func GetCartProductList(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}
}

func DeleteCartProduct(ctx *gin.Context) {

}

// cart 默认存在 使用userId 只要添加绑定记录到本表, 即说明有商品在购物车中, 否则空空如也
func GetCartxxxxxctList(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}
}

// cart 默认存在 使用userId 只要添加绑定记录到本表, 即说明有商品在购物车中, 否则空空如也
func GetCartPrx(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}
}

func GetCartPrzList(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}
}

// cart 默认存在 使用userId 只要添加绑定记录到本表, 即说明有商品在购物车中, 否则空空如也
func GetCartPrSctList(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}
}

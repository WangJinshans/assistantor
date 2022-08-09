package user

import (
	"assistantor/common"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var UserApi ApiUser

type ApiUser struct {
}

func (*ApiUser) GetUserInfo(ctx *gin.Context) {
	type req struct {
		UserId string `json:"user_id"`
	}
	var parameter req
	err := ctx.BindJSON(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}

	user, err := repository.UserRepos.GetUser(parameter.UserId)
	if err != nil {
		log.Info().Msgf("fail to get user, error is: %v", err)
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": user,
	})
	return
}

// 升级为会员必传order_id
// 取消会员不传order_id
func (*ApiUser) UpdateUserLevel(ctx *gin.Context) {
	type req struct {
		UserId  string `json:"user_id"`
		OrderId string `json:"order_id"`
	}
	var updateType int
	var parameter req
	var user *model.User
	err := ctx.BindJSON(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}
	if parameter.OrderId != "" {
		updateType = common.CreateVip // 维持vip
	} else {
		updateType = common.CancelVip // 取消vip
	}

	user, err = repository.UserRepos.GetUser(parameter.UserId)
	if err != nil {
		log.Info().Msgf("fail to get user, error is: %v", err)
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user not found",
		})
		return
	}

	var m int
	m, err = repository.UserRepos.GetUpgradeInfo(parameter.OrderId)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "order not found",
		})
		return
	}
	if updateType == common.CreateVip {
		ts := utils.GetExpireDate(m) // 过期时间
		user.LevelExpire = ts
		user.UserLevel = common.VipLevel
		if user.FirstVip == common.FirstVip {
			// 首次
			user.FirstVip = common.NotFirstVip
		}
	} else {
		user.UserLevel = common.RegularLevel
	}
	err = repository.UserRepos.UpdateUser(user)
	if err != nil {
		log.Info().Msgf("fail to get user, error is: %v", err)
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": "",
	})
	return
}

package user

import (
	"assistantor/common"
	"assistantor/repository"
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

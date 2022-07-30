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
	userId, ok := ctx.Params.Get("user_id")
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "user id not found",
		})
		return
	}

	user, err := repository.UserRepos.GetUser(userId)
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

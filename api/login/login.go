package login

import (
	"assistantor/auth"
	"assistantor/common"
	"assistantor/global"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/utils"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetPublicKey(c *gin.Context) {
	publicKeyStr, privateKeyStr, err := utils.GenRsaKey(1024)
	if err != nil {
		log.Error().Msg("密钥文件生成失败！")
		c.JSON(500, gin.H{
			"message": "server error",
		})
		return
	}
	key := fmt.Sprintf("%s%d", uuid.New().String(), time.Now().Unix())
	global.AddKeyValue(key, privateKeyStr, publicKeyStr)
	c.JSON(200, gin.H{
		"message": publicKeyStr,
		"ctx_id":  key,
	})
}

func Login(context *gin.Context) {

	var user model.User
	err := context.BindJSON(&user)
	if err != nil {
		log.Info().Msg(err.Error())
		context.JSON(200, gin.H{
			"message": "password error",
		})
		return
	}
	contextId := context.GetHeader("ctx_id")

	privateKey, err := global.GetPrivateKey(contextId)
	if err != nil {
		log.Error().Msgf("Decrypt error: %v", err)
		return
	}

	var data []byte
	data, err = base64.StdEncoding.DecodeString(user.PassWord)
	if err != nil {
		log.Error().Msgf("base64 error is: %v", err.Error())
		return
	}

	passWord, err := utils.RsaDecrypt(data, privateKey)
	if err != nil {
		log.Error().Msgf("rsa decrypt error: %v", err.Error())
		return
	}
	u, err := repository.UserRepos.GetUser(user.UserId)
	if err != nil {
		log.Error().Msgf("error is: %v", err.Error())
		return
	}
	log.Info().Msgf("password is: %s, data base password is: %s", passWord, u.PassWord)
	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), passWord) //验证（对比）
	if err != nil {
		log.Error().Msgf("error is: %v", err.Error())
		return
	}
	var token string
	token, err = auth.GenerateToken(user.UserId)
	if err != nil {
		context.JSON(200, gin.H{
			"message": "failed",
			"token":   "",
		})
		return
	}
	err = repository.SaveToken(token, user.UserId)
	if err != nil {
		context.JSON(200, gin.H{
			"message": "failed",
			"token":   "",
		})
		return
	}
	context.JSON(200, gin.H{
		"message": "success",
		"token":   token,
	})
}

func Register(context *gin.Context) {

	var user model.User
	err := context.BindJSON(&user)
	if err != nil {
		log.Error().Msg(err.Error())
		context.JSON(200, gin.H{
			"message": "password error",
		})
		return
	}

	contextId := context.GetHeader("ctx_id")

	log.Info().Msgf("encrypt password is %s", user.PassWord)
	data, err := base64.StdEncoding.DecodeString(user.PassWord)
	if err != nil {
		log.Error().Msgf("base64 decode error: %v", err.Error())
		return
	}

	privateKey, err := global.GetPrivateKey(contextId)
	if err != nil {
		log.Error().Msgf("Decrypt error: %v", err)
		return
	}

	passWord, err := utils.RsaDecrypt(data, privateKey)
	if err != nil {
		log.Error().Msgf("rsa decrypt error: %v", err.Error())
		return
	}
	userId := fmt.Sprintf("user_%d", time.Now().Unix())
	user.UserId = userId
	user.PassWord = string(passWord)
	repository.UserRepos.SaveUser(&user)
	context.JSON(200, gin.H{
		"message": userId,
	})
}

func RefreshToken(ctx *gin.Context) {
	token, err := auth.Refresh(ctx)
	if err != nil {
		log.Info().Msgf("refresh token error: %v", err)
		ctx.Abort()
		return
	}
	if token != "" {
		ctx.JSON(200, gin.H{
			"code":    common.Success,
			"message": "",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    common.EmptyToken,
			"message": "empty token",
		})
	}
	return
}

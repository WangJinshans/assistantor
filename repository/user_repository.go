package repository

import (
	"assistantor/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var UserRepos UserRepository

type UserRepository struct {
}

func (u *UserRepository) SaveUser(user *model.User) (err error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost) //加密处理
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msgf("hash string is: %s", hashStr)
	user.PassWord = string(hashStr)
	engine.Save(user)
	return
}

func (u *UserRepository) GetUser(userId string) (user model.User, err error) {
	user.UserId = userId
	engine.First(&user)
	return
}

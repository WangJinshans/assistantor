package repository

import (
	"assistantor/common"
	"assistantor/model"
	"errors"
	"github.com/matryer/try"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var UserRepos UserRepository

type UserRepository struct {
}

func (u *UserRepository) SaveUserPassword(user *model.User) (err error) {
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

func (u *UserRepository) GetUser(userId string) (user *model.User, err error) {
	user = new(model.User)
	user.UserId = userId
	engine.First(user)
	return
}

func (u *UserRepository) UpdateUser(user *model.User) (err error) {
	err = engine.Save(user).Error
	return
}

func (u *UserRepository) GetUpgradeInfo(orderId string) (total int, err error) {

	var ov []model.OrderProduct
	getFunc := func(attempt int) (bool, error) {
		ov, err = GetOrderProducts(orderId)
		if err != nil {
			log.Info().Msgf("get order products error is: %v", err)
			return attempt < 3, err
		}
		if len(ov) == 0 {
			err = errors.New(common.EmptyProductError)
			return true, err
		}

		total = int(ov[0].Count)
		return true, err // try 3 times
	}
	if err = try.Do(getFunc); err != nil {
		log.Error().Msgf("save error: %v", err)
	}
	return
}

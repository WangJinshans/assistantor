package user

import (
	"github.com/gin-gonic/gin"
)

var UserApi ApiUser

type ApiUser struct {
}

func (*ApiUser) GetUserInfo(context *gin.Context) {

}

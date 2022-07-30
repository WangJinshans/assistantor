package role

import (
	"assistantor/global"
	"github.com/gin-gonic/gin"
)

var RoleApi ApiRole

type ApiRole struct {
}

func (*ApiRole) GetAllRoles(ctx *gin.Context) {
	enforcer := global.GetEnforcer()
	roleList := enforcer.GetAllRoles()
	ctx.JSON(200, gin.H{
		"message": roleList,
	})
	return
}


func (*ApiRole) AddRoleForUser(ctx *gin.Context) {
	enforcer := global.GetEnforcer()
	userId, ok := ctx.Params.Get("user_id")
	if !ok{
		ctx.JSON(200, gin.H{
			"message": "",
		})
	}
	role, ok := ctx.Params.Get("role")
	if !ok{
		ctx.JSON(200, gin.H{
			"message": "",
		})
	}
	result, err := enforcer.AddRoleForUser(userId, role)
	if err!=nil{
		ctx.JSON(200, gin.H{
			"message": "",
		})
	}
	if !result{
		ctx.JSON(200, gin.H{
			"message": "",
		})
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
	return
}
package role

import (
	"assistantor/common"
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
		"code":    common.Success,
		"message": roleList,
	})
	return
}

func (*ApiRole) DeleteRole(ctx *gin.Context) {
	type req struct {
		Role string `json:"role"`
	}
	var parameter req
	err := ctx.BindJSON(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "empty role",
		})
		return
	}
	enforcer := global.GetEnforcer()
	ok, err := enforcer.DeleteRole(parameter.Role)
	if err != nil || !ok {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "delete role error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": "ok",
	})
	return
}

func (*ApiRole) AddRoleForUser(ctx *gin.Context) {
	enforcer := global.GetEnforcer()
	type req struct {
		UserId string `json:"user_id"`
	}
	var parameter req
	err := ctx.BindJSON(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "",
		})
		return
	}
	userId, ok := ctx.Params.Get("user_id")
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "",
		})
	}
	role, ok := ctx.Params.Get("role")
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "",
		})
	}
	result, err := enforcer.AddRoleForUser(userId, role)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "",
		})
	}
	if !result {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "",
		})
	}
	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": "success",
	})
	return
}

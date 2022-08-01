package debug

import (
	"assistantor/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"testing"
)

func TestCasbin(t *testing.T) {

	engine, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Info().Msgf("create db error: %v", err)
		return
	}

	policy, err := gormadapter.NewAdapterByDB(engine)
	if err != nil {
		log.Info().Msgf("init auth error: %v", err)
		return
	}
	enforcer, err := casbin.NewEnforcer("./auth_model.conf", policy)
	if err != nil {
		log.Info().Msgf("init auth model error: %v", err)
		return
	}
	global.SetEnforcer(enforcer) // 全局化保存
	enforcer.EnableLog(true)     // 开启权限认证日志
	// 加载数据库中的策略
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Info().Msgf("loadPolicy error: %v", err)
		panic(err)
	}

	// 创建一个角色,并赋于权限
	//admin 这个角色可以访问GET 方式访问 /api/v2/ping
	var ok bool
	ok, err = enforcer.AddPolicy("admin", "/api/v2/ping", "GET")
	if !ok {
		log.Info().Msg("policy is exist")
	} else {
		log.Info().Msg("policy is not exist, adding")
	}

	enforcer.AddRoleForUser("tom", "admin")
	enforcer.AddRoleForUser("test", "root")
	enforcer.DeleteUser("test")
	enforcer.AddRoleForUser("lance", "admin")
}

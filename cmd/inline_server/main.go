package main

import (
	"assistantor/api/login"
	"assistantor/api/v1/user"
	"assistantor/config"
	_ "assistantor/docs"
	"assistantor/global"
	"assistantor/middlerware"
	"assistantor/repository"
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 根目录: swag init -g cmd/inline_server/main.go

var (
	conf        config.AssistantConfig
	engine      *gorm.DB
	redisClient *redis.Client
)

func init() {
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	initDatabase()
	// initRedis()
	// initAuth()
}

func initDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.UserName, conf.Mysql.Password, conf.Mysql.Address, conf.Mysql.Database)

	log.Info().Msgf("connection string is: %s", dsn)
	engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	err = engine.AutoMigrate(&User{}, &Company{})
	if err != nil {
		panic(err)
	}
	repository.SetupEngine(engine)
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	res, err := redisClient.Set("movie_key", "value", time.Second*60).Result()
	log.Info().Msgf("res is: %s, error is: %v", res, err)
	repository.SetupRedisClient(redisClient)
}

func initAuth() {
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
	// admin 这个角色可以访问GET 方式访问 /api/v2/ping
	//var ok bool
	//ok,err = enforcer.AddPolicy("admin","/api/v2/ping","GET")
	//if !ok {
	//	log.Info().Msg("policy is exist")
	//} else {
	//	log.Info().Msg("policy is not exist, adding")
	//}

	//enforcer.AddRoleForUser("tom","admin")
	//enforcer.AddRoleForUser("test","root")
	//enforcer.DeleteUser("test")
	//enforcer.AddRoleForUser("lance", "admin")
}

func StartServer() {
	r := gin.Default()
	r.Use(middlerware.Cors())
	r.Use(gin.Recovery())

	r.GET("/get_public_key", login.GetPublicKey)
	r.POST("/login", login.Login)
	r.POST("/logout", login.Logout)
	r.POST("/register", login.Register)
	r.POST("/refresh_token", login.RefreshToken)
	r.GET("/get_qrcode", login.InitQrCode)
	r.GET("/get_qrcode_status", login.QueryQrCode)
	r.POST("/set_qrcode_status", login.SetQrCodeStatus)
	r.POST("/scan_qrcode", login.ScanQrCode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("v1")
	v1.Use(middlerware.JwtAuth())      // token
	v1.Use(middlerware.AuthMiddleWare) // 权限
	{
		userGroup := v1.Group("user")
		{
			userGroup.GET("/user_info", user.UserApi.GetUserInfo)
		}
	}

	r.Run(":8088")
}

type User struct {
	gorm.Model
	ID        int       `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	Name      string    `gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	Companies []Company `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
}

type Company struct {
	gorm.Model
	Industry int    `gorm:"TYPE:INT(11);DEFAULT:0"`
	Name     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX"`
	Job      string `gorm:"TYPE:VARCHAR(255);DEFAULT:''"`
	UserId   int    `gorm:"TYPE:int(11);NOT NULL;INDEX"`
}

func main() {

	p, _ := global.GetExecutablePath()
	log.Info().Msgf("p is: %s", p)
	ctx := context.Background()
	go global.StartCleanKey(ctx)
	StartServer()

}

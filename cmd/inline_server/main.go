package main

import (
	"assistantor/api/login"
	"assistantor/api/v1/user"
	"assistantor/config"
	_ "assistantor/docs"
	"assistantor/global"
	"assistantor/middlerware"
	"assistantor/model"
	"assistantor/repository"
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
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
	initRedis()
}

func initDatabase() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.UserName, conf.Mysql.Password, conf.Mysql.Address, conf.Mysql.Database)

	log.Info().Msgf("connection string is: %s", dsn)
	engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = engine.AutoMigrate(&model.User{})
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

func StartServer() {
	r := gin.Default()
	r.Use(middlerware.Cors())
	r.Use(gin.Recovery())

	r.GET("/get_public_key", login.GetPublicKey)
	r.POST("/login", login.Login)
	r.POST("/logout", login.Logout)
	r.POST("/register", login.Register)
	r.POST("/refresh_token", login.RefreshToken)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("v1")
	v1.Use(middlerware.JwtAuth())
	{
		userGroup := v1.Group("user")
		{
			userGroup.GET("/user_info", user.UserApi.GetUserInfo)
		}
	}

	r.Run(":8088")
}

func main() {

	ctx := context.Background()
	go global.StartCleanKey(ctx)
	StartServer()
}

package main

import (
	"assistantor/api/login"
	"assistantor/api/role"
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
	initRedis()
	initAuth()
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

	//err = engine.AutoMigrate(&model.User{}, &model.FilePartition{}, &model.PartitionInfo{}, &model.Message{}, &model.MediaResource{}, &model.DeliveryAddress{}, &model.Address{}, &model.Store{}, &model.Stock{}, &model.OrderProduct{}, &model.Order{})
	err = engine.AutoMigrate(&model.User{}, &model.DeliveryAddress{}, &model.Address{}, &model.Store{}, &model.Stock{}, &model.OrderProduct{}, &model.Order{})
	if err != nil {
		panic(err)
	}
	repository.SetupEngine(engine)

	CreateData()
}

func CreateData() {

	//da1 := model.DeliveryAddress{
	//	ProvinceName: "ddd",
	//	DistrictName: "dd",
	//	StreetName:   "d",
	//	AddressType:  1,
	//}
	//da2 := model.DeliveryAddress{
	//	ProvinceName: "ddd2",
	//	DistrictName: "dd2",
	//	StreetName:   "d2",
	//}
	u := model.User{
		UserLevel: 1,
		UserId:    "11111",
		UserName:  "lance",
		//AddressInfo: model.Address{
		//	ProvinceName: "aaa",
		//	DistrictName: "bbb",
		//	StreetName:   "ccc",
		//},
		//DeliveryAddressInfo: []model.DeliveryAddress{
		//	da1, da2,
		//},
	}
	//
	//err := engine.Create(&da1).Error
	//log.Info().Msgf("error is: %v", err)
	//err = engine.Create(&da2).Error
	//log.Info().Msgf("error is: %v", err)
	err := engine.Create(&u).Error
	log.Info().Msgf("error is: %v", err)

	//p1 := model.OrderProduct{
	//	ProductName: "hhhh",
	//	Description: "eat",
	//	StockInfo: model.Stock{
	//		Address:   "xxxxxxxxxxxxxxxzzzzzzzzz",
	//		StockName: "stock",
	//	},
	//}
	//
	//p2 := model.OrderProduct{
	//	ProductName: "hhxzzhh",
	//	Description: "play",
	//	StockInfo: model.Stock{
	//		Address:   "ccccccccccccccccccccccc",
	//		StockName: "stock2",
	//	},
	//}
	//
	//o := model.Order{
	//	OrderID:     "xxxxxxxxxxxxxx",
	//	OrderStatus: 1,
	//	UserInfo:    u,
	//	AddressInfo: da1,
	//	ProductList: []model.OrderProduct{
	//		p1, p2,
	//	},
	//}
	//
	//err := engine.Save(&o).Error
	//log.Info().Msgf("error is: %v", err)
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	_, err := redisClient.Set("movie_key", "value", time.Second*60).Result()
	if err != nil {
		panic(err)
	}
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
}

func StartServer() {
	r := gin.New()
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
	v1.Use(middlerware.JwtAuth())        // token
	v1.Use(middlerware.AuthMiddleWare()) // 权限
	{
		userGroup := v1.Group("user")
		{
			userGroup.GET("/user_info", user.UserApi.GetUserInfo)
		}
		roleGroup := v1.Group("role")
		{
			roleGroup.GET("/role_info", role.RoleApi.GetAllRoles)
			roleGroup.DELETE("/delete_role", role.RoleApi.DeleteRole)
			roleGroup.PUT("/add_user_role", role.RoleApi.AddRoleForUser)
		}
	}

	r.Run(":8088")
}

func main() {

	p, _ := global.GetExecutablePath()
	log.Info().Msgf("p is: %s", p)
	ctx := context.Background()
	go global.StartCleanKey(ctx)
	StartServer()

}

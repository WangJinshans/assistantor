package main

import (
	"assistantor/api/login"
	"assistantor/api/role"
	"assistantor/api/v1/cart"
	"assistantor/api/v1/order"
	"assistantor/api/v1/product"
	"assistantor/api/v1/stock"
	"assistantor/api/v1/user"
	"assistantor/config"
	_ "assistantor/docs"
	"assistantor/global"
	"assistantor/middlerware"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/services"
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
	confluentKafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 根目录: swag init -g cmd/inline_server/main.go

var (
	conf        config.AssistantConfig
	engine      *gorm.DB
	redisClient *redis.Client
	producer    *confluentKafka.Producer
)

func init() {
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	//initDatabase()
	//initRedis()
	//initAuth()
}

func initDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.UserName, conf.Mysql.Password, conf.Mysql.Address, conf.Mysql.Database)
	//log.Info().Msgf("connection string is: %s", dsn)
	engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	err = SyncTables()
	if err != nil {
		panic(err)
	}
	repository.SetupEngine(engine)
}

func SyncTables() (err error) {
	err = engine.AutoMigrate(
		&model.User{},
		&model.FilePartition{}, &model.PartitionInfo{},
		&model.Message{}, &model.MediaResource{},
		&model.Address{}, &model.DeliveryAddress{},
		&model.Store{},
		&model.Depository{},
		&model.StoreProduct{}, &model.OrderProduct{}, &model.CartProduct{},
		&model.Order{},
	)
	return
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	repository.SetupRedisClient(redisClient)
}

func initKafkaProducer() {
	var err error
	producer, err = services.NewKafkaProducer(conf.Kafka.Address)
	if err != nil {
		panic(err)
	}
	services.SetupKafkaProducer(producer)
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
	r.Static("/static", "../../static/")
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
	//v1.Use(middlerware.JwtAuth())        // token
	//v1.Use(middlerware.AuthMiddleWare()) // 权限
	{
		userGroup := v1.Group("user")
		{
			//userGroup.Use(middlerware.JwtAuth())        // token
			//userGroup.Use(middlerware.AuthMiddleWare()) // 权限
			userGroup.GET("/user_info", user.UserApi.GetUserInfo)
			userGroup.GET("/upgrade_user", user.UserApi.UpdateUserLevel)

			userGroup.PUT("/add_stock_operation", stock.AddOperationLog)
			userGroup.GET("/get_stock_operation", stock.GetOperationLog)
		}
		roleGroup := v1.Group("role")
		{
			roleGroup.GET("/role_info", role.RoleApi.GetAllRoles)
			roleGroup.DELETE("/delete_role", role.RoleApi.DeleteRole)
			roleGroup.PUT("/add_user_role", role.RoleApi.AddRoleForUser)
		}
		cartGroup := v1.Group("role")
		{
			cartGroup.GET("/cart_info", cart.GetCartProductList)
			cartGroup.DELETE("/delete_cart", cart.DeleteCartProduct)
		}
		orderGroup := v1.Group("order")
		{
			orderGroup.PUT("/create_order", order.CreateVipMemberOrder)
			orderGroup.PUT("/create_member_order", order.CreateRegularOrder)
			orderGroup.GET("/query_order", order.QueryOrderStatus)
			orderGroup.POST("/pay_order", order.PayOrder)
		}
		productGroup := v1.Group("product")
		{
			productGroup.GET("/product_list", product.GetStoreProductList)
			productGroup.GET("/product_info", product.GetStoreProductByProductId)
			productGroup.POST("/update_product", product.UpdateStoreProduct)
			productGroup.PUT("/add_product", product.AddStoreProduct)
			productGroup.DELETE("/delete_product", product.DeleteStoreProduct)
		}
		stockGroup := v1.Group("stock")
		{
			stockGroup.GET("/stock_history", stock.GetStockHistoryInfo)
			stockGroup.GET("/global_stock", stock.GetGlobalStockInfo)
			stockGroup.GET("/asia_stock", stock.GetAsiaStockInfo)
			stockGroup.GET("/amer_stock", stock.GetAmerStockInfo)
			stockGroup.GET("/europe_stock", stock.GetEuropeStockInfo)
			stockGroup.GET("/aus_stock", stock.GetAusStockInfo)
		}
	}

	r.Run(":8088")
}

func main() {

	//company_report.GetAssetsLiabilityReport("SZ002156")
	//company_report.GetProfitReport("SZ002156")
	//company_report.GetCashFlowReport("SZ002156")

	//company_report.GenerateCashFlowReport("SZ002156")
	//company_report.GenerateAssetsLiabilityReport("SZ002156")
	//company_report.GenerateProfitReport("SZ002156")

	ctx := context.Background()
	go services.StartDispatchOrder(ctx, &conf.Kafka)  // 订单分发
	go services.OrderTimeoutMonitor(ctx, &conf.Redis) // 订单超时
	go global.StartCleanKey(ctx)
	StartServer()
}

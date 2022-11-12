package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/necmettindev/currency-conversion/common/favicon"
	"github.com/necmettindev/currency-conversion/configs"
	"github.com/necmettindev/currency-conversion/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"github.com/necmettindev/currency-conversion/controllers"
	"github.com/necmettindev/currency-conversion/docs"
	"github.com/necmettindev/currency-conversion/repositories/accountrepo"
	"github.com/necmettindev/currency-conversion/repositories/userrepo"
	"github.com/necmettindev/currency-conversion/services/accountservice"
	"github.com/necmettindev/currency-conversion/services/authservice"
	"github.com/necmettindev/currency-conversion/services/financeservice"
	"github.com/necmettindev/currency-conversion/services/userservice"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	docs.SwaggerInfo.Title = "Currency Conversion Service"
	docs.SwaggerInfo.Description = "This is a currency conversion service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	if config.Env == "development" {
		router.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.SetTrustedProxies(nil)

	db, err := gorm.Open(
		config.Postgres.Dialect(),
		config.Postgres.GetPostgresConnectionInfo(),
	)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	redisConfig := configs.GetRedisConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	userRepo := userrepo.NewUserRepo(db)
	accountRepo := accountrepo.NewAccountRepo(db)

	userService := userservice.NewUserService(userRepo, config.Pepper)
	authService := authservice.NewAuthService(config.JWTSecret)
	financeService := financeservice.NewFinanceService()
	accountService := accountservice.NewAccountService(rdb, accountRepo)

	userCtl := controllers.NewUserController(userService, authService)
	accountCtl := controllers.NewAccountController(userService, accountService, financeService)

	router.Use(favicon.New("./public/favicon.ico"))
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Currency Conversion Service", "status": http.StatusOK})
	})

	api := router.Group("/v1")

	user := api.Group("/users")
	user.POST("/register", userCtl.PostRegister)
	user.POST("/login", userCtl.PostLogin)

	account := api.Group("/accounts")
	account.Use(middlewares.RequireLoggedIn(config.JWTSecret))
	{
		account.GET("/", accountCtl.GetBalances)
		account.GET("/:first_currency/:second_currency/rate", accountCtl.GetCurrencyConversionRate)
		account.POST("/:first_currency/:second_currency/:amount/conversion", accountCtl.PostCurrencyConversion)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + config.Port)
}

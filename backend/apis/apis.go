package apis

import (
	"fmt"
	"time"

	"sirichai-bank/drivers"
	"sirichai-bank/dtos"

	_ "sirichai-bank/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title sirichai-bank API
// @version 1.0.0
// @description This is a used car listing application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// var (
// 	appVersion = "v1.0.0"
// 	apiVersion = "api/v1"
// 	Fiber      = fiber.Config{
// 		ServerHeader: "sirichai-bank" + appVersion,
// 		BodyLimit:    10 * 1024 * 1024,
// 	}
// 	FiberCore = cors.New(cors.Config{
// 		AllowOrigins: "*",
// 		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
// 		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
// 	})
// )

// func NewFiberAPI(signKey string, mongoDB *drivers.MongoDBClient, srvAuth services.AuthService, srvList services.ListService, srvMsg services.MessageService) *fiber.App {
// 	f := fiber.New(Fiber)
// 	f.Use(FiberCore)
// 	f.Use(recover.New())

// 	f.Use(func(c *fiber.Ctx) error {
// 		c.Locals(dtos.KEY_SECRET, signKey)

// 		log.Printf("Request: %s %s from %s", c.Method(), c.Path(), c.IP())
// 		return c.Next()
// 	})

// 	f.Group(apiVersion).Get("/swagger/*", fiberSwagger.WrapHandler)

// 	NewHandlerAuth(f, srvAuth, signKey)
// 	NewHandlerList(f, srvList, srvAuth)
// 	NewHandlerMessage(f, srvList, srvAuth, srvMsg)

// 	return f
// }

func NewContext(code int, msg string, res any) *dtos.Context {
	return &dtos.Context{
		Status:  code,
		Message: msg,
		Results: res,
	}
}

func NewGinAPI(signKey string, firebase *drivers.FirebaseClient) *gin.Engine {
	g := gin.Default()

	// Configure CORS
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	g.Use(gin.Recovery())
	g.Use(func(c *gin.Context) {
		c.Set(dtos.KEY_SECRET, signKey)
		c.Next()
	})

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "OK",
			"results": "sirichai-bank API is working",
		})
	})

	g.POST("/api/v1/verifytoken", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "Unauthorized",
				"results": "Token is required",
			})
		}

		TokenValid,err := firebase.Auth.VerifyIDToken(c, token)
		if err != nil {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "Unauthorized",
				"results": "Token is invalid",
			})
		}

		fmt.Printf("TokenValid: %+v\n", TokenValid)
		firebase.DB.Collection("users").Add(c, map[string]interface{}{
			"uid": TokenValid.UID,
			"email": "test@mail.com",
			"created_at": time.Now(),
			"updated_at": time.Now(),
		})


		c.JSON(200, gin.H{
			"status":  200,
			"message": "OK",
			"results": "Token is valid",
		})
	})

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// NewHandlerAuthGin(r, srvAuth, signKey)
	// NewHandlerListGin(r, srvList, srvAuth)
	// NewHandlerMessageGin(r, srvList, srvAuth, srvMsg)

	return g
}

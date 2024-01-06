package routes

import (
	"API/login"
	"API/user_management_admin"
	"API/AREA_management"
	"API/about"

	"github.com/gin-contrib/cors"
	"github.com/swaggo/files"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
)

func Routes() {
	router := gin.Default()


	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://front-area.lekmax.fr", "http://localhost:8081", "exp://*.kapwiing.8081.exp.direct", "http://localhost:8080", "http://localhost:8000","https://front-area.lekmax.fr"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "x-api-key")
	config.AllowHeaders = append(config.AllowHeaders, "bearer")
	router.Use(cors.New(config))

	router.Static("/docs", "./docs")
	url := ginSwagger.URL("/docs/swagger.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/user", user_management_admin.GetUser)
	router.POST("/login", login.Login)
	router.POST("/register", login.Register)
	router.POST("/create-area", AREA_management.PostArea)
	router.POST("/modify-area", AREA_management.PostModifyArea)
	router.GET("/areas", AREA_management.CheckArea)
	router.POST("/get-area", AREA_management.PostGetAllArea)
	router.POST("/login/github", login.GithubLogin)
	router.POST("/delete-area", AREA_management.PostDeleteArea)
	router.GET("/about.json", about.DetailedAbout)
	router.POST("/login/discord", login.DiscordLogin)

	router.Run("0.0.0.0:8080")
}

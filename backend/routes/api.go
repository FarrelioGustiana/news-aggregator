package routes

import (
	controllers "github.com/FarrelioGustiana/backend/controllers"
	middleware "github.com/FarrelioGustiana/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.RegisterUser)
		authRoutes.POST("/login", controllers.LoginUser)
	}

	apiRoutes := router.Group("/api")

	apiRoutes.Use(middleware.AuthMiddleware())
	{
		// Feeds
		apiRoutes.POST("/feeds", controllers.CreateFeed)
		apiRoutes.GET("/feeds", controllers.GetAllFeeds)
		apiRoutes.GET("/feeds/:id", controllers.GetFeedByID)
		apiRoutes.PUT("/feeds/:id", controllers.UpdateFeed)
		apiRoutes.DELETE("/feeds/:id", controllers.DeleteFeed)

		// Subcriptions
		apiRoutes.POST("/subscriptions", controllers.SubscribeToFeed) 
		apiRoutes.GET("/subscriptions", controllers.GetUserSubscriptions)
		apiRoutes.DELETE("/subscriptions/:feedId", controllers.UnsubscribeFromFeed) 

	}
}

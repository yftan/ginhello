package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	router.Static("/statics", "./statics")

	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)

		userRouter.POST("/login",handler.UserLogin)

		userRouter.GET("/:name", handler.UserSave)

		userRouter.GET("", handler.UserSaveQuery)
	}

	return router
}


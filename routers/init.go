package routers

import (
	handlers "yuanaturego/handlers/frontend"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//首頁(關於我們)
	router.GET("/", handlers.About)
	router.GET("/about", handlers.About)

	return router
}

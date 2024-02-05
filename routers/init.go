package routers

import (
	//handlersAdmin "yuanaturego/handlers/backend"
	handlers "yuanaturego/handlers/frontend"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//首頁(關於我們)
	router.GET("/", handlers.About)
	router.GET("/about", handlers.About)
	//購買商品
	router.GET("/product", handlers.Product)
	//訂單
	/*ordersRouter := router.Group("/orders")
	{
		//購物車
		ordersRouter.GET("/cart", handlers.OrdersCart)
	}

	//後台
	adminRouter := router.Group("/admin")
	{
		//首頁
		adminRouter.GET("/", handlersAdmin.Index)
	}

	//AJAX
	ajaxRouter := router.Group("/ajax")
	{
		ajaxRouter.POST("/cart_data", handlers.AjaxCart)
	}*/

	return router
}

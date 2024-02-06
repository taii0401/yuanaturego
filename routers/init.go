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
	//購物須知
	router.GET("/shopping", handlers.Shopping)
	//運送政策
	router.GET("/shipment", handlers.Shipment)
	//退換貨政策
	router.GET("/refunds", handlers.Refunds)
	//隱私權政策
	router.GET("/privacy", handlers.Privacy)
	//購物問題
	router.GET("/qa_shopping", handlers.QaShopping)
	//產品問題
	router.GET("/qa_product", handlers.QaProduct)
	//會員問題
	router.GET("/qa_member", handlers.QaMember)
	//使用者回饋
	//router.GET("/feedback", handlers.Feedback)
	//router.GET("/feedback_detail", handlers.FeedbackDetail)
	//聯絡我們
	//router.GET("/contact", handlers.Contact)

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

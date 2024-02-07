package frontend

import (
	"net/http"
	handlers "yuanaturego/handlers"
	models "yuanaturego/models"

	"github.com/gin-gonic/gin"
)

// 關於我們
func About(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "關於我們"
	context.HTML(http.StatusOK, "about.tmpl", data)
}

// 購買商品
func Product(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "購買商品"

	//取得商品資料
	product := models.ProductModel{
		Uuid: "8c8f7d49-2131-4a50-a02d-619f3e1f9fb9",
	}
	product_data := product.GetDataByUuid()
	for key, value := range product_data {
		data[key] = value
	}

	//商品說明圖片
	img := make(map[int]int)
	for i := 1; i <= 13; i++ {
		img[i] = i
	}
	data["img"] = img
	context.HTML(http.StatusOK, "product.tmpl", data)
}

// 購物須知
func Shopping(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "購物須知"
	context.HTML(http.StatusOK, "shopping.tmpl", data)
}

// 運送政策
func Shipment(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "運送政策"
	context.HTML(http.StatusOK, "shipment.tmpl", data)
}

// 退換貨政策
func Refunds(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "退換貨政策"
	context.HTML(http.StatusOK, "refunds.tmpl", data)
}

// 隱私權政策
func Privacy(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "隱私權政策"
	context.HTML(http.StatusOK, "privacy.tmpl", data)
}

// 購物問題
func QaShopping(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "購物問題"
	context.HTML(http.StatusOK, "qa_shopping.tmpl", data)
}

// 產品問題
func QaProduct(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "產品問題"
	context.HTML(http.StatusOK, "qa_product.tmpl", data)
}

// 會員問題
func QaMember(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "會員問題"
	context.HTML(http.StatusOK, "qa_member.tmpl", data)
}

//使用者回饋

// 聯絡我們
func Contact(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "聯絡我們"
	//選項-聯絡我們類型
	data["contact_type"] = handlers.GetConfigOptions("contact_type", true)
	context.HTML(http.StatusOK, "contact.tmpl", data)
}

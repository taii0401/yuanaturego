package frontend

import (
	"net/http"
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

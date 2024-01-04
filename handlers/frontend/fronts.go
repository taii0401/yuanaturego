package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 關於我們
func About(context *gin.Context) {
	data := make(map[string]interface{})
	data["title_txt"] = "關於我們"
	context.String(http.StatusOK, "About")
	//context.HTML(http.StatusOK, "about.tmpl", data)
}

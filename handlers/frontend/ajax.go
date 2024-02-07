package frontend

import (
	"net/http"
	models "yuanaturego/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var error = true
var message = "資料錯誤"

// 聯絡我們-新增
func AjaxContactData(context *gin.Context) {
	//動作類型
	action_type := context.PostForm("action_type")

	//驗證欄位

	contact := models.ContactModel{}
	if action_type == "add" { //新增
		uuid := uuid.New().String()
		contact.Uuid = uuid
		id := -1
		if err := context.ShouldBind(&contact); err == nil {
			id = contact.Create()
		}
		if id > 0 {
			error = false
			message = uuid
		} else {
			message = "送出失敗"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"error":   error,
		"message": message,
	})
}

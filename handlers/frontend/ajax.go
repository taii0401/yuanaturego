package frontend

import (
	"net/http"
	handlers "yuanaturego/handlers"
	models "yuanaturego/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var error = true
var message = "資料錯誤"

// 聯絡我們-新增
func AjaxContactData(context *gin.Context) {
	//動作類型
	action_type := context.PostForm("action_type")

	//驗證欄位
	validField := make(map[string]string)
	validField["Name"] = "姓名"
	validField["Phone"] = "電話"

	contact := models.ContactModel{}
	if action_type == "add" { //新增
		uuid := uuid.New().String()
		contact.Uuid = uuid

		errMsg := ""
		id := -1
		if err := context.ShouldBind(&contact); err == nil {
			id = contact.Create()
		} else {
			errMsg = handlers.ValidateMsg(err.(validator.ValidationErrors), validField)
		}

		if id > 0 {
			error = false
			message = uuid
		} else {
			message = errMsg
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"error":   error,
		"message": message,
	})
}

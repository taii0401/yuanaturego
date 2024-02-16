package handlers

import (
	"github.com/go-playground/validator/v10"
)

var configDatas = map[string]map[string]interface{}{
	//動作名稱
	"action_name": {"add": "新增", "edit": "編輯", "delete": "刪除", "cancel": "取消", "import": "匯入", "export": "匯出", "search": "查詢", "other": "其他"},
	//登入後轉向頁面
	"login_url":      {"/product": nil},
	"login_url_cart": {"/orders/cart": nil},
	//登出後轉向頁面
	"logout_url": {"/users": nil},
	//綠界超商類型
	"ecpay_store": {"UNIMARTC2C": "seven", "FAMIC2C": "family", "OKMARTC2C": "ok", "HILIFEC2C": "hilife"},
}

var configOptions = map[string]map[int]map[string]interface{}{
	//登入方式
	"user_register": {
		0: {"key": "email", "name": "EMAIL", "color": "#696969"},
		1: {"key": "facebook", "name": "FACEBOOK", "color": "#2626FF"},
		2: {"key": "line", "name": "LINE", "color": "#008F00"},
		3: {"key": "google", "name": "GOOGLE", "color": "#FF2626"},
	},

	//訂單狀態
	"orders_status": {
		0: {"key": "nopaid", "name": "未付款", "color": "#B87800"},
		1: {"key": "paid", "name": "已付款", "color": "#8100E0"},
		2: {"key": "failpaid", "name": "付款失敗", "color": "#E00000"},
		3: {"key": "refund", "name": "已退款", "color": "#FF2626"},
		4: {"key": "handle", "name": "處理中", "color": "#696969"},
		5: {"key": "cancel", "name": "已取消", "color": "#E00000"},
		6: {"key": "delivery", "name": "已出貨", "color": "#008F00"},
		7: {"key": "back", "name": "已退貨", "color": "#E00000"},
		8: {"key": "complete", "name": "已完成", "color": "#008A8A"},
	},

	//付款方式
	"orders_payment": {
		0: {"key": "card", "name": "信用卡", "color": "#B87800"},
		1: {"key": "atm", "name": "ATM", "color": "#696969"},
		2: {"key": "linepay", "name": "LINE PAY", "color": "#008F00"},
	},

	//配送方式
	"orders_delivery": {
		0: {"key": "home", "name": "宅配配送", "color": "#008A8A"},
		1: {"key": "store", "name": "超商取貨", "color": "#B87800"},
	},

	//台灣本島或離島
	"orders_island": {
		0: {"key": "main", "name": "台灣本島", "color": "#B87800"},
		1: {"key": "outlying", "name": "台灣離島", "color": "#008A8A"},
	},

	//取消原因
	"orders_cancel": {
		0: {"key": "wrong", "name": "下錯訂單", "color": "#E00000"},
		1: {"key": "forget", "name": "忘記使用優惠", "color": "#B87800"},
		2: {"key": "unbuy", "name": "不想購買", "color": "#B87800"},
		3: {"key": "other", "name": "其他", "color": "#00B800"},
		4: {"key": "system", "name": "系統", "color": "#696969"},
	},

	//超商類型
	"orders_store": {
		0: {"key": "seven", "name": "7-11", "color": "#B87800"},
		1: {"key": "family", "name": "全家", "color": "#8100E0"},
		2: {"key": "ok", "name": "OK", "color": "#696969"},
		3: {"key": "hilife", "name": "萊爾富", "color": "#E00000"},
	},

	//聯絡我們-問題類型
	"contact_type": {
		0: {"key": "user", "name": "會員相關問題", "color": "#B87800"},
		1: {"key": "product", "name": "商品相關問題", "color": "#8100E0"},
		2: {"key": "order", "name": "訂單相關問題", "color": "#008A8A"},
		3: {"key": "pay", "name": "付款相關問題", "color": "#FF2626"},
		4: {"key": "delivery", "name": "運費、寄送問題", "color": "#696969"},
		5: {"key": "refund", "name": "退貨、退款問題", "color": "#E00000"},
		6: {"key": "other", "name": "其他", "color": "#008F00"},
	},

	//聯絡我們-處理狀態
	"contact_status": {
		0: {"key": "handle", "name": "處理中", "color": "#696969"},
		1: {"key": "complete", "name": "已處理", "color": "#008A8A"},
		2: {"key": "cancel", "name": "已取消", "color": "#E00000"},
	},

	//折價劵-使用狀態
	"coupon_status": {
		0: {"key": "nouse", "name": "未使用", "color": "#696969"},
		1: {"key": "used", "name": "已使用", "color": "#008A8A"},
		2: {"key": "expire", "name": "已過期", "color": "#FF2626"},
		3: {"key": "cancel", "name": "已取消", "color": "#E00000"},
	},

	//折價劵-發放來源
	"coupon_source": {
		0: {"key": "register", "name": "會員註冊", "color": "#696969"},
		1: {"key": "admin", "name": "管理員贈送", "color": "#008A8A"},
	},
}

/**
* 依類型取得設定檔選項
* @param  config_type：設定檔類型
* @param  is_all：是否回傳全部的選項
* @return map
 */
func GetConfigOptions(config_type string, is_all bool) map[int]map[string]string {
	data := make(map[int]map[string]string)
	i := 0
	if is_all {
		data[i] = map[string]string{
			"key":  "",
			"name": "全部",
		}
		i++
	}

	for j := 0; j <= len(configOptions[config_type]); j++ {
		key := configOptions[config_type][j]["key"].(string)
		name := configOptions[config_type][j]["name"].(string)

		//訂單取消原因，移除系統選項
		if config_type == "orders_cancel" && key == "system" {
			continue
		}

		data[i] = map[string]string{
			"key":  key,
			"name": name,
		}
		i++
	}

	return data
}

/**
* 驗證欄位訊息
* @param  error：錯誤訊息
* @param  validField：欄位名稱
* @return string
 */
func ValidateMsg(err validator.ValidationErrors, validField map[string]string) string {
	result := make(map[string]string, 0)
	i := 0
	for _, v := range err {
		if field, ok := v.(validator.FieldError); ok {
			name := field.Field()
			tag := field.Tag()
			msg := validField[name]

			if tag == "required" {
				msg += "為必填"
			} else {
				msg += "格式錯誤"
			}

			result[string(i)] = msg
			i++
		}
	}

	return MapToString(result)
}

// 將map轉換成字串
func MapToString(data map[string]string) string {
	str := ""
	i := 0
	for _, value := range data {
		str += value
		if i < len(data)-1 {
			str += "，"
		}
		i++
	}

	return str
}

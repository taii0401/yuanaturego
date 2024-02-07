package handlers

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

var configOptions = map[string]map[string]map[string]interface{}{
	//登入方式
	"user_register": {
		"email":    {"name": "EMAIL", "color": "#696969"},
		"facebook": {"name": "FACEBOOK", "color": "#2626FF"},
		"line":     {"name": "LINE", "color": "#008F00"},
		"google":   {"name": "GOOGLE", "color": "#FF2626"},
	},

	//訂單狀態
	"orders_status": {
		"nopaid":   {"name": "未付款", "color": "#B87800"},
		"paid":     {"name": "已付款", "color": "#8100E0"},
		"failpaid": {"name": "付款失敗", "color": "#E00000"},
		"refund":   {"name": "已退款", "color": "#FF2626"},
		"handle":   {"name": "處理中", "color": "#696969"},
		"cancel":   {"name": "已取消", "color": "#E00000"},
		"delivery": {"name": "已出貨", "color": "#008F00"},
		"back":     {"name": "已退貨", "color": "#E00000"},
		"complete": {"name": "已完成", "color": "#008A8A"},
	},

	//付款方式
	"orders_payment": {
		"card":    {"name": "信用卡", "color": "#B87800"},
		"atm":     {"name": "ATM", "color": "#696969"},
		"linepay": {"name": "LINE PAY", "color": "#008F00"},
	},

	//配送方式
	"orders_delivery": {
		"home":  {"name": "宅配配送", "color": "#008A8A"},
		"store": {"name": "超商取貨", "color": "#B87800"},
	},

	//台灣本島或離島
	"orders_island": {
		"main":     {"name": "台灣本島", "color": "#B87800"},
		"outlying": {"name": "台灣離島", "color": "#008A8A"},
	},

	//取消原因
	"orders_cancel": {
		"wrong":  {"name": "下錯訂單", "color": "#E00000"},
		"forget": {"name": "忘記使用優惠", "color": "#B87800"},
		"unbuy":  {"name": "不想購買", "color": "#B87800"},
		"other":  {"name": "其他", "color": "#00B800"},
		"system": {"name": "系統", "color": "#696969"},
	},

	//超商類型
	"orders_store": {
		"seven":  {"name": "7-11", "color": "#B87800", "logistic": "UNIMARTC2C"},
		"family": {"name": "全家", "color": "#8100E0", "logistic": "FAMIC2C"},
		"ok":     {"name": "OK", "color": "#696969", "logistic": "OKMARTC2C"},
		"hilife": {"name": "萊爾富", "color": "#E00000", "logistic": "HILIFEC2C"},
	},

	//聯絡我們-問題類型
	"contact_type": {
		"user":     {"name": "會員相關問題", "color": "#B87800"},
		"product":  {"name": "商品相關問題", "color": "#8100E0"},
		"order":    {"name": "訂單相關問題", "color": "#008A8A"},
		"pay":      {"name": "付款相關問題", "color": "#FF2626"},
		"delivery": {"name": "運費、寄送問題", "color": "#696969"},
		"refund":   {"name": "退貨、退款問題", "color": "#E00000"},
		"other":    {"name": "其他", "color": "#008F00"},
	},

	//聯絡我們-處理狀態
	"contact_status": {
		"handle":   {"name": "處理中", "color": "#696969"},
		"complete": {"name": "已處理", "color": "#008A8A"},
		"cancel":   {"name": "已取消", "color": "#E00000"},
	},

	//折價劵-使用狀態
	"coupon_status": {
		"nouse":  {"name": "未使用", "color": "#696969"},
		"used":   {"name": "已使用", "color": "#008A8A"},
		"expire": {"name": "已過期", "color": "#FF2626"},
		"cancel": {"name": "已取消", "color": "#E00000"},
	},

	//折價劵-發放來源
	"coupon_source": {
		"register": {"name": "會員註冊", "color": "#696969"},
		"admin":    {"name": "管理員贈送", "color": "#008A8A"},
	},
}

/**
* 依類型取得設定檔選項
* @param  cond：搜尋條件
* @param  is_all：是否回傳全部的選項
* @return array
 */
func GetConfigOptions(config_type string, is_all bool) map[string]interface{} {
	data := make(map[string]interface{})
	if is_all {
		data[""] = "全部"
	}
	for key, value := range configOptions[config_type] {
		//訂單取消原因，移除系統選項
		if config_type == "orders_cancel" && key == "system" {
			continue
		}
		data[key] = value["name"]
	}

	return data
}

package models

import (
	"reflect"
	"time"
	database "yuanaturego/database"
)

type ProductModel struct {
	Id         uint      `gorm:"column:id;primaryKey;autoIncrement;not null;"`
	Uuid       string    `gorm:"column:uuid;unique;not null;type:uuid;default:uuid_generate_v4();"`
	Category   string    `gorm:"column:category"`
	SerialCode string    `gorm:"column:serial_code"`
	SerialNum  int       `gorm:"column:serial_num"`
	Serial     string    `gorm:"column:serial"`
	Name       string    `gorm:"column:name"`
	Price      int       `gorm:"column:price"`
	Sales      int       `gorm:"column:sales"`
	Content    string    `gorm:"column:content;default:null"`
	IsDisplay  int       `gorm:"column:is_display;default:1"`
	CreatedId  int       `gorm:"column:created_id;default:null"`
	UpdatedId  int       `gorm:"column:updated_id;default:null"`
	DeletedId  int       `gorm:"column:deleted_id;default:null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  time.Time `gorm:"column:deleted_at;default:null"`
}

func (ProductModel) TableName() string {
	return "product"
}

// 依訂單UUID取得資料
func (model_data ProductModel) GetDataByUuid() map[string]interface{} {
	database.Db.Where("deleted_at IS NULL").Find(&model_data, ProductModel{
		Uuid: model_data.Uuid,
	}).Rows()

	//轉換map
	data := model_data.changeFieldName()

	return data
}

// 轉換map
func (model_data ProductModel) changeFieldName() map[string]interface{} {
	data := make(map[string]interface{})
	//取得欄位名稱
	columns, _ := database.Db.Migrator().ColumnTypes(ProductModel{})

	field_data := reflect.ValueOf(&model_data).Elem()
	for i := 0; i < field_data.NumField(); i++ {
		//欄位名稱
		field_name := columns[i].Name()
		//欄位值
		field_value := field_data.Field(i).Interface()
		data[field_name] = field_value
	}

	return data
}

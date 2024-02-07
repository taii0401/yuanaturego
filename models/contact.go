package models

import (
	"reflect"
	"time"
	database "yuanaturego/database"
)

type ContactModel struct {
	Id        uint      `gorm:"column:id;primaryKey;autoIncrement;not null;"`
	Uuid      string    `gorm:"column:uuid;unique;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"column:name" form:"name"`
	Phone     string    `gorm:"column:phone" form:"phone"`
	Email     string    `gorm:"column:email" form:"email"`
	Type      string    `gorm:"column:type;default:null" form:"type"`
	Status    string    `gorm:"column:status;default:handle" form:"status"`
	Content   string    `gorm:"column:content;default:null" form:"content"`
	Reply     string    `gorm:"column:reply;default:null" form:"reply"`
	CreatedId int       `gorm:"column:created_id;default:null"`
	UpdatedId int       `gorm:"column:updated_id;default:null"`
	DeletedId int       `gorm:"column:deleted_id;default:null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:null"`
}

func (ContactModel) TableName() string {
	return "contact"
}

// 新增資料
func (model_data ContactModel) Create() int {
	result := database.Db.Create(&model_data)
	if result.Error != nil {
		return 0
	}
	return int(model_data.Id)
}

// 依UUID取得資料
func (model_data ContactModel) GetDataByUuid() map[string]interface{} {
	database.Db.Where("deleted_at IS NULL").Find(&model_data, ContactModel{
		Uuid: model_data.Uuid,
	}).Rows()

	//轉換資料
	data := model_data.changeFieldName()

	return data
}

// 轉換資料
func (model_data ContactModel) changeFieldName() map[string]interface{} {
	data := make(map[string]interface{})
	//取得欄位名稱
	columns, _ := database.Db.Migrator().ColumnTypes(ContactModel{})

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

package main

import (
	"encoding/json"
	oracle "github.com/dzwvip/gorm-oracle"
	"log"
	"strings"
	"time"

	_ "github.com/sijms/go-ora/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Sex_dict struct {
	SerialNo int    `json:"serial_no" db:"SERIAL_NO" gorm:"not null;primaryKey;column:SERIAL_NO;"`
	SexCode  string `json:"Sex_code" db:"SEX_CODE" gorm:"not null;column:SEX_CODE;"`
	SexName  string `json:"Sex_name" db:"SEX_NAME" gorm:"column:SEX_NAME;"`
}

func (*Sex_dict) TableName() string {
	return "comm.sex_dict"
}

type Language struct {
	Code      string    `json:"code" gorm:"not null;primaryKey;type:varchar(2);comment:语言代码"`
	Name      string    `json:"name" gorm:"not null;type:varchar(60)"`
	Lang      string    `json:"lang" gorm:"type:varchar(60);DEFAULT:en;"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime:milli"`
}

func (*Language) TableName() string {
	return "COMM.LANGUAGES"
}

type CustomReplacer struct{}

func (r CustomReplacer) Replace(name string) string {
	return strings.ToUpper(name)
}

func main() {
	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("gorm测试 开始Oracle Driver example")
	dsn := "oracle://system:manager@localhost:1521/zmhis"
	//dsn := "oracle://system:manager@192.168.56.102:1521/this"
	ns := schema.NamingStrategy{
		NameReplacer:  CustomReplacer{},
		NoLowerCase:   true,
		SingularTable: true,
	}

	gormdb, err := gorm.Open(oracle.Open(dsn), &gorm.Config{NamingStrategy: ns}) //

	// err := gormdb.Debug().AutoMigrate(
	// 	Language{},
	// )
	if err != nil {
		log.Println(err)
	}

	// 原生sql语句
	dicts2 := []Sex_dict{}

	// _ = gormdb.Debug().Model(&Sex_dict{}).Where("SERIAL_NO = ?", 7).Limit(1).Find(&dicts2).Error
	//_ = gormdb.Debug().Model(&Sex_dict{}).Limit(1).Find(&dicts2).Error
	_ = gormdb.Debug().First(&dicts2).Error
	jsonData2, err := json.Marshal(dicts2)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonData2))

	//dicts1 := Sex_dict{
	//	SerialNo: 6,
	//
	//	SexCode: "6",
	//	SexName: "测6",
	//}
	//gormdb.Debug().Save(&dicts1)

	gdb, _ := gormdb.DB()

	_ = gdb.Close()

}

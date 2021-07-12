package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//为了项目使用方便，有些变量可以定义为全局。
var (
	DB *gorm.DB
)

//链接mysql
func InitMySQL() (err error) {
	dsn := "root:12321@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}

func InitModel() {
	DB.AutoMigrate()
}

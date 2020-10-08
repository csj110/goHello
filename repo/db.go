package repo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySql() (err error) {
	dsn := "root:sqlsecret@tcp(127.0.0.1:3306)/goHello?charset=utf8&parseTime=true&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}
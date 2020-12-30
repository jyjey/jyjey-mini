package mysql

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB 数据库链接
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/jyjey?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}

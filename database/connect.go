package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func New() *gorm.DB {
	db, sqlerr := sqlConnect()
	if sqlerr != nil {
		panic(sqlerr.Error())
	} else {
		fmt.Println("DB接続成功")
	}

	return db
}


// func New() (*gorm.DB, error) {
// 	db, err := sqlConnect()
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	fmt.Println("DB接続成功")

// 	return db, nil
// }


// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "Naka3Naka4"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_practice"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

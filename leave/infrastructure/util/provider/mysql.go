package provider

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type DataBase struct {
	Instance *gorm.DB
}

func getDBInstance(address, username, password string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		address,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (db *DataBase) NewDataBase(address, username, password string) *DataBase {
	return &DataBase{
		Instance: getDBInstance(address, username, password),
	}
}

func (db *DataBase) Close() {
	if err := db.Instance.Close(); err != nil {
		log.Fatal("close db error: ", err)
	}
}

package config

import (
	"fmt"

	"github.com/aungkoko1234/gin-crud/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 3306
	user = "root"
	password = "root"
	dbName ="go_crud"
)

func DatabaseConnection() *gorm.DB {	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password, host, port, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
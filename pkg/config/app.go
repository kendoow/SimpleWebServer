package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db * gorm.DB
)


func Connect(){
	d, err := gorm.Open("mysql", "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local") // парольи логин своей бд убрать при пуше.
	if err != nil{ 
		panic(err)
	}
	db = d
}

func GetBD() *gorm.DB  {
	return db
}
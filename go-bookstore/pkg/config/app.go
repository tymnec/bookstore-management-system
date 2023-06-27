// App

package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect Function
func Connect() {
	d, err := gorm.Open("mysql", "tymnec:Nikhilsarwara@123/simplerest?charset=utf8&parseTime=True&loc=Local")

	// Check if there is an error
	if err != nil {
		panic(err)
	}

	db = d
}

// GetDB Function
func GetDB() *gorm.DB {
	return db
}

package database

import (
	"github.com/ttvs-blockchain/firebaseupdate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(conf *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(conf.MySQLdsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

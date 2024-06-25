package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/userManagement?parseTime=true" // MySQL credentials and database name
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func CloseDB() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"todo-api/models"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=pia password=1234 dbname=todo_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	time.Sleep(10 * time.Second)

	if err != nil {
		panic("Faild to connect to Database.")
	}

	db.AutoMigrate(&models.Task{})
	return db
}

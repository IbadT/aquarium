package database

import (
	"aquarium/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "host=localhost user=postgres password=postgres dbname=aquarium port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Ошибка подключения к базе данных")
	}

	log.Println("Подключение к базе данных")

	DB.AutoMigrate(&models.User{})

	log.Println("Таблицы успешно созданы")
}

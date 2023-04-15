package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"todo-list/mvc/model"
)

var DB *gorm.DB

func ConfigInit() {

	initEnv()
	initDatabase()

}

func initEnv() {

	err := godotenv.Load(RootPath() + `\.env.test`)
	if err != nil {
		ENV_STATUS = ERROR
		log.Fatal(">> Error loading .env file")
	} else {
		ENV_STATUS = SUCCESS
	}

}

func initDatabase() {

	var dsn string

	if ENV_STATUS == SUCCESS {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USERNAME")
		pwd := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_DATABASE")

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dbName)

	} else {
		dsn = "test:test@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		DB_STATUS = ERROR
		log.Fatal(">> Failed connect database")

	} else {
		DB_STATUS = SUCCESS
		DB = db
		err := db.AutoMigrate(&model.User{}, &model.Todo{})

		if err != nil {
			log.Fatal(">> Failed create tables")
		}
	}

}

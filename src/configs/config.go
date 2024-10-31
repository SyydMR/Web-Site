package configs

import (
	"fmt"
	"log"

	"github.com/SyydMR/Web-Site/src/handlers"
	"github.com/SyydMR/Web-Site/src/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
    DB   *gorm.DB
)


func loadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./src/configs")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
}

func connectDB() {
	loadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.port"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	
	DB = db
	log.Println("Database connection established successfully.")

}


func Init() *gorm.DB {
	connectDB()
	DB.AutoMigrate(&models.User{}, &models.Task{}, &models.Content{}, &models.Post{})
	models.InitDB(DB)
	handlers.InitDB(DB)
	return DB
}
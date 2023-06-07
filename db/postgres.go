package db

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"log"
	"github.com/ShivamIITK21/cflockout-backend/models"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SslMode  string
}

func GenConfig() *Config {
	config := &Config{
		Host: os.Getenv("HOST"),
		User: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
		SslMode: os.Getenv("SSL_MODE"),
	}
	return config
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting to Database")
	}
	return db, err
}

func AutoMigrateAll(db *gorm.DB) error{
	err := db.AutoMigrate(&models.Problem{})
	if err != nil {
		log.Fatal("Error during Automigrate")
	}
	return err
}
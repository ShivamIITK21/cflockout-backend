package db

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"log"
	"github.com/ShivamIITK21/cflockout-backend/models"
	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SslMode  string
}

func genConfig() *Config {
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

func newConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func autoMigrateAll(db *gorm.DB) error{
	err := db.AutoMigrate(&models.Problem{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Lockout{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.User{})
	return err
}

func setupDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	config := genConfig()
	db, err := newConnection(config)
	if err != nil {
		log.Fatal("Error Connecting to Postgres")
	}
	err = autoMigrateAll(db)
	if err != nil {
		log.Fatal("Error in Migrating Models")
	}
	return db
}

var DB = setupDB()
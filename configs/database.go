package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, using system environment variables")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set in environment")
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect database: ", err)
	}

	// Test connection (ping)
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("❌ Failed to get DB instance: ", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("❌ Database unreachable: ", err)
	}

	log.Println("✅ Database connected successfully!")
	DB = database
}

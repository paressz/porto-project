package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"porto-project/models"
)

type Database struct {
	Instance *gorm.DB
}

var Db Database
func loadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed loading .env")
	}
	return os.Getenv(key)
}
func Connect() {
	var host, user, password, port, dbname = loadEnv("PGHOST"), loadEnv("PGUSER"), loadEnv("PGPASSWORD"), loadEnv("PGPORT"), loadEnv("PGDBNAME")
	dsn :=
		"host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: \n", err)
	}
	log.Print("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Project{})
	Db = Database{ Instance: db }
}

package PDB

import (
	"log"
	model "lyked-backend/model/posgressModels"
	"lyked-backend/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func ConnectPostgres() (*gorm.DB, error) {
	utils.LoadEnv()
	dsn := utils.GetEnv("POSTGRES_CONNECTION_STRING", "BACKUP_POSTGRES_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("✅ Connected to PostgreSQL database")
	if err := db.AutoMigrate(&model.User{}, &model.Session{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("✅ Database migrated")
	PostgresDB = db
	return db, nil
}

package migration

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func RunMigration() {
	log.Println("Start Migration Database ...")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	addr := os.Getenv("DB_HOSTNAME")
	dbname := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@tcp(" + addr + ")/" + dbname
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database M					igrated")

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()
}

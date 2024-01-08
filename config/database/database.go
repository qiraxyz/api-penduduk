package database

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	config := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOSTNAME"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	DB, err := sql.Open(dbDriver, config.FormatDSN())
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Second)
	DB.SetConnMaxLifetime(60 * time.Minute)

	err = pingDb(DB)
	if err != nil {
		return DB, err
	}

	return DB, nil
}

func DBConnectionDw() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	config := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOSTNAME"),
		DBName:               os.Getenv("DB_NAME_DW"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	DB, err := sql.Open(dbDriver, config.FormatDSN())
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Second)
	DB.SetConnMaxLifetime(60 * time.Minute)
	err = pingDb(DB)
	if err != nil {
		return DB, err
	}

	return DB, nil
}

func DBConnectionDwAgg() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	config := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOSTNAME"),
		DBName:               os.Getenv("DB_NAME_DW_AGG"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	DB, err := sql.Open(dbDriver, config.FormatDSN())
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Second)
	DB.SetConnMaxLifetime(60 * time.Minute)
	err = pingDb(DB)
	if err != nil {
		return DB, err
	}

	return DB, nil
}

func pingDb(db *sql.DB) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

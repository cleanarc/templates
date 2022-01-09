package persistence

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DB struct {
	*gorm.DB
}

func (d *DB) GetConn() (*sql.DB, error) {
	conn, err := d.DB.DB()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Setup() *DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=America/Detroit",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	testConn(db)

	err = migrate(db)
	if err != nil {
		return nil
	}

	return &DB{db}
}

func testConn(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	// Ping the database to make sure its up and healthy
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Pinged DB successfully")
}

func migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(&db_models.Model{})
	if err != nil {
		return err
	}

	return nil
}

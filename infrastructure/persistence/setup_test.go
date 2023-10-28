package persistence

import (
	"fmt"
	"gadget-points/domain/entity"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBConn() (*gorm.DB, error) {
	if _, err := os.Stat("./../../.env"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("./../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return LocalDatabase()
	}
	return CIBuild()
}

// Circle CI DB
func CIBuild() (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "steven", "gadget-points-test", "password")
	conn, err := gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return conn, nil
}

// Local DB
func LocalDatabase() (*gorm.DB, error) {
	dbdriver := os.Getenv("TEST_DB_DRIVER")
	host := os.Getenv("TEST_DB_HOST")
	password := os.Getenv("TEST_DB_PASSWORD")
	user := os.Getenv("TEST_DB_USER")
	dbname := os.Getenv("TEST_DB_NAME")
	port := os.Getenv("TEST_DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	conn, err := gorm.Open(dbdriver, DBURL)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", dbdriver)
	}

	err = conn.DropTableIfExists(&entity.User{}, &entity.Agent{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		entity.User{},
		entity.Agent{},
	).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

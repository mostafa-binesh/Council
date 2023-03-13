package database

import (
	"fmt"
	"log"
	"os"
	"time"

	U "docker/utils"
	// "gorm.io/driver/mysql"
	// U "docker/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDatabase *gorm.DB

func DB() *gorm.DB {
	if gormDatabase != nil {
		return gormDatabase
	}
	ConnectToDB()
	return gormDatabase
}

func ConnectToDB() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	DB_SERVER, _ := U.Env("DB_SERVER") // localhost name and port
	DB_USERNAME, _ := U.Env("DB_USERNAME")
	DB_PASSWORD, _ := U.Env("DB_PASSWORD")
	DB_NAME, _ := U.Env("DB_NAME") // database name
	DB_PORT, _ := U.Env("DB_PORT")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USERNAME, DB_PASSWORD, DB_SERVER, DB_PORT, DB_NAME)
	gormDatabase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Printf("database connection stablished.")
}
func RowsCount(query string, searchValue string) int {
	rows, err := gormDatabase.Raw(query, searchValue).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
		break
	}

	return count
}

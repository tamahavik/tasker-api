package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second,   // Slow SQL threshold
	//		LogLevel:                  logger.Silent, // Log level
	//		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	//		Colorful:                  false,         // Disable color
	//	},
	//)

	connectionStr := `postgres://postgres:password@localhost:5432/tasker?sslmode=disable`
	sqlDriver, err1 := sql.Open("postgres", connectionStr)
	if err1 != nil {
		return nil, err1
	}

	sqlDriver.SetMaxIdleConns(10)
	sqlDriver.SetMaxOpenConns(100)
	sqlDriver.SetConnMaxIdleTime(5 * time.Minute)
	sqlDriver.SetConnMaxLifetime(60 * time.Minute)

	db, err2 := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDriver,
	}), &gorm.Config{})
	if err2 != nil {
		return nil, err2
	}

	return db, nil

}

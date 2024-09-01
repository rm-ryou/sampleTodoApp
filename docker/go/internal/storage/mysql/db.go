package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	db              *sql.DB
	MaxDBRetryCount = 10
)

func migrateDB() error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprint("file:///migrations"),
		"mysql", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func connectDB(driver, dsn string) error {
	var err error

	for i := 1; i <= MaxDBRetryCount; i++ {
		db, err = sql.Open(driver, dsn)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}

		if err = db.Ping(); err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	db.SetConnMaxIdleTime(time.Second * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}

func SetUpDB(driver, dsn string) (*sql.DB, error) {
	if err := connectDB(driver, dsn); err != nil {
		return nil, err
	}
	if err := migrateDB(); err != nil {
		return nil, err
	}

	return db, nil
}

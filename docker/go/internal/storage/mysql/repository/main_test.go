package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	pwd, _   = os.Getwd()
	testDB   *sql.DB
	ur       *UserRepository
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

func createDB() error {
	var err error

	pool, err = dockertest.NewPool("")
	pool.MaxWait = time.Minute
	if err != nil {
		return err
	}

	err = pool.Client.Ping()
	if err != nil {
		return err
	}

	resource, err = pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=root",
			"MYSQL_DATABASE=test_db",
			"MYSQL_USER=user",
			"MYSQL_PASSWORD=password",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return err
	}

	return nil
}

func connectDB() error {
	if err := pool.Retry(func() error {
		var err error

		testDB, err = sql.Open("mysql", fmt.Sprintf("user:password@tcp(localhost:%s)/test_db?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return testDB.Ping()
	}); err != nil {
		return err
	}

	return nil
}

func migrateDB() error {
	driver, err := mysql.WithInstance(testDB, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/../migrations", pwd),
		"mysql", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func setUp() error {
	if err := createDB(); err != nil {
		return err
	}
	if err := connectDB(); err != nil {
		return err
	}
	if err := migrateDB(); err != nil {
		return err
	}

	return nil
}

func teardown() {
	if err := pool.Purge(resource); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func TestMain(m *testing.M) {
	if err := setUp(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	ur = NewUserRepository(testDB)

	m.Run()
	teardown()
}

package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go/modules/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
)

const (
	dbUsername string = "root"
	dbPassword string = "root"
	dbName     string = "test"
)

var (
	db *sqlx.DB
)

func SetupMySQLContainer() (func(), *sqlx.DB, error) {
	log.Println("setup MySQL Container")
	ctx := context.Background()

	seedDataPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("%v", err))
	}

	mysqlC, err := mysql.RunContainer(ctx,
		testcontainers.WithImage("mysql:latest"),
		mysql.WithDatabase(dbName),
		mysql.WithUsername(dbUsername),
		mysql.WithPassword(dbPassword),
		mysql.WithScripts(filepath.Join(seedDataPath, "/../../test", "data.sql")),
	)

	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("%v", err))
	}

	closeContainer := func() {
		log.Println("terminating container")
		err := mysqlC.Terminate(ctx)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
	}

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Int()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify&parseTime=true&multiStatements=true",
		dbUsername, dbPassword, host, port, dbName)

	db, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		log.Println(err)
		return closeContainer, db, err
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		return closeContainer, db, err
	}

	return closeContainer, db, nil
}

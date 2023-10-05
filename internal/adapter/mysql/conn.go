package mysql

import (
	"fmt"
	"log"

	"github.com/hieronimusbudi/simple-go-api/internal/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (db *sqlx.DB) {
	host := config.Get().DBHOST
	user := config.Get().DBUSER
	pass := config.Get().DBPASSWORD
	dbName := config.Get().DBNAME
	descriptor := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbName)

	db, err := sqlx.Connect("mysql", descriptor)
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return db
}

package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewDb(dbHost string, dbPort int, dbName, dbUser, dbPass string) *sqlx.DB {
	connStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Open("mysql", connStr)
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		log.WithError(err).Fatalf("db connect error (%s)", connStr)
	}

	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxOpenConns(500)
	return db
}

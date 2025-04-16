package main

import (
	"fmt"
	"os"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetMysqlClient(msUser, msPassword, msHost, msPort, msDatabase string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		msUser, msPassword, msHost, msPort, msDatabase)
	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(msUser, msPassword, msHost, msPort, msDatabase string) error {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?multiStatements=true",
		msUser, msPassword, msHost, msPort, msDatabase)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	if err != nil {
		return err
	}

	wd, _ := os.Getwd()
	path := "file://" + wd + "/database/"

	m, _ := migrate.NewWithDatabaseInstance(
		path,
		"mysql",
		driver,
	)

	m.Up()

	return nil
}

package utils

import (
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetMysqlClient() *gorm.DB {
	dsn := "root:1234@(db:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate() {
	db, err := sql.Open("mysql", "root:1234@(db:3306)/todo?multiStatements=true")
	if err != nil {
		panic(err)
	}

	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	m.Steps(2)
}

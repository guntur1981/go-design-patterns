package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// a facade
type MySQLDB struct {
	db *sql.DB
}

func (m *MySQLDB) Connect() error {
	mysqlcfg := mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%S", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		DBName:               os.Getenv("MYSQL_DBNAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
	}

	db, err := sql.Open("mysql", mysqlcfg.FormatDSN())
	if err != nil {
		return err
	}

	m.db = db
	return nil
}

func (m *MySQLDB) Ping() bool {
	err := m.db.Ping()
	return err == nil
}

func main() {
	mysqlDb := MySQLDB{}
	err := mysqlDb.Connect()
	if err != nil {
		log.Fatalf("Error connecting MySQL: %s\n", err)
	}

	mysqlDb.Ping()
}

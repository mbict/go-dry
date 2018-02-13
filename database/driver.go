package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/mattn/go-sqlite3"
)

const (
	MYSQL    = "mysql"
	POSTGRES = "postgres"
	SQLITE3  = "sqlite3"
)

// DriverName will return the dialect driver name
func DriverName(db *sql.DB) string {
	switch db.Driver().(type) {
	case *sqlite3.SQLiteDriver:
		return SQLITE3
	case *pq.Driver:
		return POSTGRES
	case *mysql.MySQLDriver:
		return MYSQL
	}
	return ""
}

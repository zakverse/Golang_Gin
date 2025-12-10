package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
    dsn := "root:DZAKI2006@tcp(127.0.0.1:3306)/db_penyewaan?parseTime=true"
    return sql.Open("mysql", dsn)
}
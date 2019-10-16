package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/opentracing-contrib/sql"
)

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{}, sql.WithDSNParser(ParseDSN))
}

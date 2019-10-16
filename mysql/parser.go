package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/opentracing-contrib/sql"
)

// ParseDSN parses the given go-sql-driver/mysql datasource name.
func ParseDSN(name string) sql.DSNInfo {
	cfg, err := mysql.ParseDSN(name)
	if err != nil {
		// mysql.Open will fail with the same error,
		// so just return a zero value.
		return sql.DSNInfo{}
	}
	return sql.DSNInfo{
		Database: cfg.DBName,
		User:     cfg.User,
	}
}

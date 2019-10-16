package pq

import (
	pq "github.com/lib/pq"
	"github.com/opentracing-contrib/sql"
)

func init() {
	sql.Register("postgres", &pq.Driver{}, sql.WithDSNParser(ParseDSN))
}

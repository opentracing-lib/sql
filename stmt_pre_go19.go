package sql

import "database/sql/driver"

type stmtGo19 struct{}

func (stmtGo19) init(in driver.Stmt) {}

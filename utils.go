package sql

import (
	"database/sql/driver"
	"errors"
)

// namedValueToValue copied from database/sql (see NOTICE).
func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			return nil, errors.New("sql: driver does not support the use of Named Parameters")
		}
		dargs[n] = param.Value
	}
	return dargs, nil
}

// namedValueChecker is identical to driver.NamedValueChecker, existing
// for compatibility with Go 1.8.
type namedValueChecker interface {
	CheckNamedValue(*driver.NamedValue) error
}

func checkNamedValue(nv *driver.NamedValue, next namedValueChecker) error {
	if next != nil {
		return next.CheckNamedValue(nv)
	}
	return driver.ErrSkip
}

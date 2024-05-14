package transaction

import "github.com/jmoiron/sqlx"

type DbSession struct {
	Ext      sqlx.ExtContext
	Commit   func() error
	Rollback func() error
}

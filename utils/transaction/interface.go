package transaction

import "github.com/jmoiron/sqlx"

type DbSession struct {
	Ext       sqlx.ExtContext
	Stmt      func(*sqlx.Stmt) *sqlx.Stmt
	NamedStmt func(*sqlx.NamedStmt) *sqlx.NamedStmt
	Commit    func() error
	Rollback  func() error
}

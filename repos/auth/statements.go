package auth

import (
	"github.com/JesseNicholas00/HaloSuster/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	createUser *sqlx.NamedStmt
	findByNip  *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		createUser: statementutil.MustPrepareNamed(`
			INSERT INTO users(
				user_id,
				nip,
				name,
				password,
				active,
				image_url
			) VALUES (
				:user_id,
				:nip,
				:name,
				:password,
				:active,
				:image_url
			) RETURNING
				user_id,
				nip,
				name,
				password,
				active,
				image_url,
				created_at
		`),
		findByNip: statementutil.MustPrepare(`
			SELECT
				*
			FROM
				users
			WHERE
				nip = $1
		`),
	}
}

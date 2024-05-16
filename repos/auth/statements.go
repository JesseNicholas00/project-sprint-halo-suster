package auth

import (
	"github.com/JesseNicholas00/HaloSuster/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	createUser            *sqlx.NamedStmt
	findByNip             *sqlx.Stmt
	activateNurseByUserId *sqlx.NamedStmt
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
		activateNurseByUserId: statementutil.MustPrepareNamed(`
			UPDATE users
			SET
				active = :active,
				password = :password
			WHERE
				user_id = :user_id AND nip < 6150000000000
			RETURNING
				user_id,
				nip,
				name,
				password,
				active,
				image_url,
				created_at
		`),
	}
}

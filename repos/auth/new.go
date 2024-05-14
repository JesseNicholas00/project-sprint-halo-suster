package auth

import (
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
)

type authRepostioryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func NewAuthRepository(dbRizzer ctxrizz.DbContextRizzer) AuthRepository {
	return &authRepostioryImpl{
		dbRizzer: dbRizzer,
	}
}

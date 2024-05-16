package auth

import (
	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
)

type authServiceImpl struct {
	repo       auth.AuthRepository
	jwtSecret  []byte
	bcryptCost int
	dbRizz     ctxrizz.DbContextRizzer
}

func NewAuthService(
	repo auth.AuthRepository,
	jwtSecret string,
	bcryptCost int,
	dbRizz ctxrizz.DbContextRizzer,
) AuthService {
	return &authServiceImpl{
		repo:       repo,
		jwtSecret:  []byte(jwtSecret),
		bcryptCost: bcryptCost,
		dbRizz:     dbRizz,
	}
}

package auth_test

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
	"github.com/JesseNicholas00/HaloSuster/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) auth.AuthRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return auth.NewAuthRepository(ctxrizz.NewDbContextRizzer(db))
}

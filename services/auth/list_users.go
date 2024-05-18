package auth

import (
	"context"
	"time"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (svc *authServiceImpl) ListUsers(
	ctx context.Context,
	req ListUsersReq,
	res *ListUsersRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	repoRes, err := svc.repo.ListUsers(ctx, auth.UserFilter{
		UserId:        req.UserId,
		Limit:         *req.Limit,
		Offset:        *req.Offset,
		Name:          req.Name,
		Nip:           req.Nip,
		Role:          req.Role,
		CreatedAtSort: req.CreatedAtSort,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, user := range repoRes {
		res.Data = append(res.Data, ListUsersResData{
			UserId:    user.Id,
			Nip:       user.Nip,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.Format(time.RFC3339Nano),
		})
	}

	return nil
}

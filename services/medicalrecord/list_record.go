package medicalrecord

import (
	"context"
	"errors"
	"time"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/transaction"
)

// ListRecord implements MedicalRecordService.
func (svc *medicalRecordServiceImpl) ListRecord(
	ctx context.Context,
	req ListRecordReq,
	res *ListRecordRes,
) (err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {

		filterUser := make(map[string]auth.User)
		if req.CreatedByNip != nil {
			userRes, err := svc.repoAuth.FindUserByNip(ctx, *req.CreatedByNip)
			if err != nil {
				if errors.Is(err, auth.ErrNipNotFound) {
					return nil
				}
				return errorutil.AddCurrentContext(err)
			}
			if req.CreatedByUserId != nil &&
				*req.CreatedByUserId != userRes.Id {
				return nil
			}
			req.CreatedByUserId = &userRes.Id
			filterUser[userRes.Id] = userRes
		}

		var (
			limit  = 5
			offset = 0
		)
		if req.Limit != nil {
			limit = *req.Limit
		}
		if req.Offset != nil {
			offset = *req.Offset
		}

		recordRes, err := svc.repo.ListRecord(ctx, medicalrecord.RecordFilter{
			IdentityNumber:  req.IdentityNumber,
			CreatedByUserId: req.CreatedByUserId,
			Limit:           limit,
			Offset:          offset,
			CreatedAtSort:   req.CreatedAtSort,
		})
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		if len(recordRes) > 0 {
			if len(filterUser) == 0 {
				userIds := getUserIds(recordRes)
				recordUsers, err := svc.repoAuth.ListAllUsers(
					ctx,
					auth.AllUsersFilter{
						UserIds: userIds,
					},
				)
				for _, recordUser := range recordUsers {
					filterUser[recordUser.Id] = recordUser
				}
				if err != nil {
					return errorutil.AddCurrentContext(err)
				}

			}
			mapToRes(recordRes, filterUser, res)
		}

		return nil

	})
}

func getUserIds(results []medicalrecord.Record) (userIds []string) {
	for _, result := range results {
		userIds = append(userIds, result.CreatedByUserId)
	}
	return userIds
}

func mapToRes(
	results []medicalrecord.Record,
	users map[string]auth.User,
	res *ListRecordRes,
) {
	var mappedResult []ListRecordResData
	for _, result := range results {
		mappedResult = append(mappedResult, ListRecordResData{
			IdentityDetail: RecordIdentityDetail{
				IdentityNumber: result.PatientId,
				PhoneNumber:    result.PatientPhoneNumber,
				Name:           result.PatientName,
				BirthDate: result.PatientBirthDate.Format(
					time.DateOnly,
				),
				Gender:              result.PatientGender,
				IdentityCardScanImg: result.PatientImageUrl,
			},
			Symptoms:    result.Symptoms,
			Medications: result.Medications,
			CreatedAt:   result.CreatedAt.Format(time.RFC3339Nano),
			CreatedBy: RecordCreatedByDetail{
				Nip:    users[result.CreatedByUserId].Nip,
				Name:   users[result.CreatedByUserId].Name,
				UserId: result.CreatedByUserId,
			},
		})
	}
	res.Data = mappedResult
}

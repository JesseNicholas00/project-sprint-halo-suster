package ctxrizz

import (
	"context"
	"database/sql"

	"github.com/JesseNicholas00/HaloSuster/utils/transaction"
	"github.com/jmoiron/sqlx"
)

type txContextKeyType struct{}

func (c txContextKeyType) String() string {
	return "txHelper"
}

var txContextKey = txContextKeyType{}

var noop = func() error {
	return nil
}

func (cb *dbContextRizzerImpl) AppendTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	tx, err := cb.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return ctx, transaction.DbSession{}, err
	}

	return context.WithValue(ctx, txContextKey, tx), transaction.DbSession{
		Ext:      tx,
		Commit:   tx.Commit,
		Rollback: tx.Rollback,
	}, nil
}

func (cb *dbContextRizzerImpl) GetOrAppendTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	// already has tx
	if tx, ok := ctx.Value(txContextKey).(*sqlx.Tx); ok {
		// commit and rollback should be done on the creating layer
		// noop here
		return ctx, transaction.DbSession{
			Ext:      tx,
			Commit:   noop,
			Rollback: noop,
		}, nil
	}

	return cb.AppendTx(ctx)
}

func (cb *dbContextRizzerImpl) GetOrNoTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	// already has tx
	if tx, ok := ctx.Value(txContextKey).(*sqlx.Tx); ok {
		// commit and rollback should be done on the creating layer
		// noop here
		return ctx, transaction.DbSession{
			Ext:      tx,
			Commit:   noop,
			Rollback: noop,
		}, nil
	}

	return ctx, transaction.DbSession{
		Ext:      cb.db,
		Commit:   noop,
		Rollback: noop,
	}, nil
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: transfers.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers(
    from_account_id,
    to_account_id, 
    amount
) VALUES ( 
    $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :one
DELETE from transfers  WHERE id = $1 RETURNING id, from_account_id, to_account_id, amount, created_at
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, deleteTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
 WHERE id = $1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE  transfers  SET amount = $2 WHERE id = $1 RETURNING id, from_account_id, to_account_id, amount, created_at
`

type UpdateTransferParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer, arg.ID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

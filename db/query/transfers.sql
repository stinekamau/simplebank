
CREATE TABLE transfers (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- name: CreateTransfer: one 
INSERT INTO transfers(
    fromAccountID,
    toAccountID, 
    Amount
) VALUES (  +
    $1, $2, $3
) RETURNING * 

-- name: GetTransfer: one
SELECT FROM transfers where id=$1;

-- name: UpdateTransfer: one 
UPDATE  transfers  SET Amount = $2 WHERE id = $1 RETURNING *;

-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions(
  id CHAR(27) NOT NULL PRIMARY KEY,
  source_account_id TEXT NOT NULL REFERENCES accounts(account_id),
  destination_account_id  TEXT NOT NULL REFERENCES accounts(account_id),
  amount NUMERIC NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd

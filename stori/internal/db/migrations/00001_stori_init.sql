-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
  "id" VARCHAR NOT NULL DEFAULT gen_random_uuid()::varchar,
  "bank_name" VARCHAR NOT NULL,
  "number" INTEGER NOT NULL,
  "currency" VARCHAR NOT NULL,
  "account_name" VARCHAR NOT NULL,
  "account_email" VARCHAR NOT NULL,
  "create_ts" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id),
  UNIQUE(account_email, number)
);

CREATE TABLE IF NOT EXISTS transaction (
  "id" VARCHAR NOT NULL DEFAULT gen_random_uuid()::varchar,
  "account_id" VARCHAR NOT NULL,
  "date" TIMESTAMP NOT NULL,
  "debit_amount" NUMERIC(10,2) DEFAULT 0.0,
  "credit_amount" NUMERIC(10,2) DEFAULT 0.0,
  "metadata" JSON,
  "create_ts" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id),
  CONSTRAINT fk_account
    FOREIGN KEY(account_id)
    REFERENCES account(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction;
DROP TABLE account;
-- +goose StatementEnd

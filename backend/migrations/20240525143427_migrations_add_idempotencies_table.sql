-- +goose Up
-- +goose StatementBegin
CREATE TABLE idempotencies(
    id CHAR(27) NOT NULL PRIMARY KEY,
    key_hash TEXT NOT NULL, 
    http_response_code INTEGER NOT NULL,
    http_response_headers JSONB NOT NULL,
    http_response_body TEXT NOT NULL,
    created_at TIMESTAMP with time zone NOT NULL DEFAULT now(),
    updated_at TIMESTAMP with time zone NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP with time zone,
    UNIQUE (key_hash)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS idempotencies;
-- +goose StatementEnd

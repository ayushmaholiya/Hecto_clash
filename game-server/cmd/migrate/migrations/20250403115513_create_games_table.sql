-- +goose Up
CREATE TABLE IF NOT EXISTS games (
    id BIGSERIAL PRIMARY KEY,
    hectoc_puzzle CHAR(6) NOT NULL CHECK (hectoc_puzzle ~ '^[0-9]{6}$'),
    winner_id BIGINT,
    winning_submission TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now()
);
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS games;
-- +goose StatementBegin
-- +goose StatementEnd

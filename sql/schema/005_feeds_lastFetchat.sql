-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetch_at TIMESTAMP NOT NULL DEFAULT now();

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetch_at;

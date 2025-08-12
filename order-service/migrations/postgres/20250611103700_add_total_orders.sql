-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    ADD COLUMN total INT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    DROP COLUMN total;
-- +goose StatementEnd

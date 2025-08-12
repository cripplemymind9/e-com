-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id          SERIAL PRIMARY KEY,
    user_id     INT         NOT NULL,
    product_id  INT         NOT NULL,
    quantity    INT         NOT NULL,
    status      VARCHAR(50) NOT NULL    DEFAULT 'processing',
    created_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd

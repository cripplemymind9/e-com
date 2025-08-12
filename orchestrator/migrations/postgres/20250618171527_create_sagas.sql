-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sagas (
    id SERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    steps JSONB NOT NULL DEFAULT '[]',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX idx_sagas_order_id ON sagas(order_id);
CREATE INDEX idx_sagas_status ON sagas(status);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sagas;
-- +goose StatementEnd

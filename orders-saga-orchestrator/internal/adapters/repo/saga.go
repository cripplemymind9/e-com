package repo

import (
	"context"
	"encoding/json"

	"github.com/cripplemymind9/orders-saga-orchestrator/internal/domain/entity"
	"github.com/jackc/pgx/v5/pgtype"
)

func (q *queries) GetSagaByOrderID(ctx context.Context, id int64) (entity.Saga, error) {
	const query = `
		SELECT
			id,
			order_id,
			status,
			steps,
			created_at,
			updated_at
		FROM sagas
		WHERE order_id = $1
	`

	row := q.db.QueryRow(ctx, query, id)

	var (
		out       entity.Saga
		stepsJSON []byte
		updatedAt pgtype.Timestamp
	)

	err := row.Scan(
		&out.ID,
		&out.OrderID,
		&out.Status,
		&stepsJSON,
		&out.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return entity.Saga{}, err
	}

	if updatedAt.Valid {
		out.UpdatedAt = &updatedAt.Time
	}

	if unmarshalErr := json.Unmarshal(stepsJSON, &out.Steps); unmarshalErr != nil {
		return entity.Saga{}, unmarshalErr
	}

	return out, nil
}

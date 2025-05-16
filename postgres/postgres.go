package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (l *DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}
func (l *DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	fmt.Printf("[SQL] %s\n", query)
	return nil
}

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}

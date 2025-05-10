package healthcheck

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type interactor struct {
	postgres *sqlx.DB
}

// NewUsecase --
func NewUsecase(postgres *sqlx.DB) Inport {
	return interactor{
		postgres: postgres,
	}
}

func (i interactor) Execute(ctx context.Context) (InportResponse, error) {
	if err := i.postgres.PingContext(ctx); err != nil {
		return InportResponse{
			Status: map[string]string{"database": "project-x database connection error"},
		}, err
	}

	return InportResponse{
		Status: map[string]string{
			"database": "healthy",
		},
	}, nil
}

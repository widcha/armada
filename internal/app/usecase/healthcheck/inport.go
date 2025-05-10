package healthcheck

import (
	"context"
)

//go:generate mockgen -destination=mock/inport.go -package=mock farm/internal/app/usecases/health_check Inport
type Inport interface {
	Execute(context.Context) (InportResponse, error)
}

type InportResponse struct {
	Status map[string]string
}

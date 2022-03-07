package reception

import (
	"admin/reception/internal/model"
	"context"
)

type Service interface {
	Get(ctx context.Context, filters ...model.Filter) ([]model.Reception, error)
	Status(ctx context.Context, receptionID string) (model.Status, error)
	AddReception(ctx context.Context, reception *model.Reception) (string, error)
	ChangeStatus(ctx context.Context, receptionID string, status string) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}

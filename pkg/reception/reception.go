package reception

import (
	"admin/reception/internal/model"
	"context"
)

type receptionService struct{}

func NewService() Service {
	return &receptionService{}
}

func (r receptionService) Get(_ context.Context, filters ...model.Filter) ([]model.Reception, error) {
	//TODO implement me
	panic("implement me")
}

func (r receptionService) Status(ctx context.Context, receptionID string) (model.Status, error) {
	//TODO implement me
	panic("implement me")
}

func (r receptionService) AddReception(ctx context.Context, reception *model.Reception) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r receptionService) ChangeStatus(ctx context.Context, receptionID string, status string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r receptionService) ServiceStatus(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

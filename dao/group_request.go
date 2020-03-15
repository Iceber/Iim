package dao

import (
	"context"
	"github.com/iceber/iim/model"
)

func (d *dao) GroupRequest(ctx context.Context, requestID string) (*model.GroupRequest, error) {
}

func (d *dao) GroupRequests(ctx context.Context, userID string) ([]model.GroupRequest, error) {
}

func (d *dao) AddGroupRequest(ctx context.Context, request *model.GroupRequest) (*model.GroupRequest, error) {
}

func (d *dao) ApproveGroupRequest(ctx context.Context, request *model.GroupRequest) error {
}

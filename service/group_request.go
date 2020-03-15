package service

import (
	"context"

	"github.com/iceber/iim/model"
)

type groupRequestService struct{}

var GroupRequestService = new(groupRequestService)

func (grs *groupRequestService) List(ctx context.Context, userID string) ([]model.GroupRequest, error) {
	requests, err := dao.GroupRequests(ctx, userID)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (grs *groupRequestService) Add(ctx context.Context, groupID, userID string) error {
	err := dao.CreateGroupRequest(ctx, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (grs *groupRequestService) Approve(ctx context.Context, requestID string) error {
	request, err := dao.GroupRequest(ctx, requestID)
	if err != nil {
		return err
	}
	return grs.Join(ctx, request.GroupID, request.UserID)
}

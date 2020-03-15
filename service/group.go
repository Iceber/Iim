package service

import (
	"context"

	"github.com/iceber/iim/model"
)

type groupService struct{}

var GroupService = new(groupService)

func (gs *groupService) List(ctx context.Context, userID string) ([]model.Group, error) {
	groups, err := dao.UserGroups(userID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (gs *groupService) Get(ctx context.Context, groupID string) (*model.Group, error) {
	groupUser, err := dao.GroupUser(groupID, userID)
	if err != nil {
		return nil, err
	}
	if groupUser == nil {
		return nil, errors.New("sdfdfdf")
	}

	group, err := dao.Group(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return group
}

func (gs *groupService) Add(ctx context.Context, creatorID string, group *model.Group) (*model.Group, error) {
	group.CreatorID = creatorID

	return dao.CreapGroup(ctx, group)
}

func (gs *groupService) Update(ctx context.Context, userID string, group *model.Group) (*model.Group, error) {
	group, err := dao.Group(ctx, group.GroupID)
	if err != nil {
		return nil, err
	}
	if group.OwnerID != userID {
		return nil, err
	}
	return dao.UpdateGroup(ctx, group)

}

func (gs *groupService) Delete(ctx context.Context, userID, groupID string) error {
	group, err := dao.Group(ctx, groupID)
	if err != nil {
		return err
	}
	if group.Owner != userID {
		return err
	}
	return dao.DeleteGroup(ctx, groupID)
}

func (gs *groupService) Join(ctx context.Context, groupID, userID string) error {
	_, err := dao.AddGroupUser(groupID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (gs *groupService) Quit(ctx context.Context, groupID, userID string) error {
	_, err := dao.DeleteGroupUser(group, userID)
	if err != nil {
		return err
	}
	return nil
}

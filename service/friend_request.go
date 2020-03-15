package service

import (
	"context"
	"errors"
	"time"

	"github.com/iceber/iim/model"
)

type friendRequestService struct{}

var FriendRequestService = new(friendRequestService)

type FriendRequest struct {
	UserInfo  UserInfo
	Status    string
	ExpiredAt time.Time
}

func (fs *friendRequestService) List(ctx context.Context, userID string) ([]model.FriendRequest, error) {
	modelRequests, err := dao.FriendRequests(ctx, userID)
	if err != nil {
		return nil, err
	}

	var userIDs []string
	for mr := range modelRequests {
		userIDs = append(userIDs, mr.FriendID)
	}
	users, err := dao.GetUserByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]model.User)
	for _, user := range users {
		userMap[user.UserID] = user
	}

	requests := make([]FriendRequest, 0, len(modelRequests))
	for mr := range modelRequests {
		r := FriendRequest{
			Status:    mr.Status,
			ExpiredAt: mr.ExpiredAt,
			UserInfo: UserInfo{
				Nickname: userMap[r.FriendID].Nickname,
			},
		}
		requests = append(requests, r)
	}
	return requests, nil
}

func (fs *friendRequestService) Add(ctx context.Context, userID, friendID string) (*model.FriendRequest, error) {
	if isFriend, err := dao.IsFriend(ctx, userID, friendID); err != nil {
		return nil, err
	} else if isFriend {
		return nil, errors.New("is friend")
	}

	request, err := dao.AddFriendRequests(ctx, userID, friendID)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (fs *friendRequestService) Approve(ctx context.Context, userID, friendID string, info *FriendInfo) error {
	requests, err := dao.FriendRequest(ctx, userID, friendID)
	if err != nil {
		return err
	}
	for r := range requests {
		r.Status = model.FriendRequestStatusApproved
		dao.UpdateFriendRequestStatus(ctx, r)
	}

	modelFriend := &model.Friend{
		UserID:   userID,
		FriendID: friendID,
		Info:     FriendInfo,
	}
	dao.AddFriend(ctx, modelFriend)
}

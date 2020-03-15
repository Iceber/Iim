package service

import (
	"context"
)

type friendService struct{}

var FriendService = new(friendService)

type Friend struct {
	UserID    string
	Nickname  string
	AvatarURL string
	Info      FriendInfo
}

type FriendInfo struct {
	RemarkName string
}

func (fs *friendService) List(ctx context.Context, userID string) ([]Friend, error) {
	modelFriends, err := dao.Friends(ctx, userID)
	if err != nil {
		return nil, err
	}
	return modelFriends, nil
}

func (fs *friendService) Get(ctx context.Context, userID, friendID string) (*Friend, error) {
	mf, err := dao.Friend(ctx, userID, friendID)
	if err != nil {
		return nil, err
	}
	return mf, nil
}

func (fs *friendService) Delete(ctx context.Context, userID, friendID string) error {
	if err := dao.DeleteFriend(ctx, userID, friendID); err != nil {
		return err
	}
	return nil
}

func (fs *friendService) Update(ctx context.Context, userID, friendID string, info *FriendInfo) (*Friend, error) {
	mf, err := dao.Friend(ctx, userID, friendID)
	if err != nil {
		return nil, err
	}

	mf.Info.RemarkName = info.RemarkName
	if err := dao.UpdateFriend(ctx, mf); err != nil {
		return nil, err
	}
	return &Friend{
		UserID:    mf.FriendID,
		Nickname:  mf.Info.Nickname,
		Info:      FriendInfo{RemarkName: mf.Info.RemarkName},
		AvatarURL: mf.Info.AvatarURL,
	}, nil
}

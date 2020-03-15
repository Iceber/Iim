package model

import "time"

type FriendInfo struct {
	AvatarURL  string
	Nickname   string
	RemarkName string `json:"remark_name"`
}

type Friend struct {
	ID        int64      `json:"-" db:"id"`
	UserID    string     `json:"user_id"`
	FriendID  string     `json:"friend_id"`
	Info      FriendInfo `json:"friend_user_info"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

const (
	FriendRequestStatusPending  = "applying"
	FriendRequestStatusApproved = "approved"
)

type FriendRequest struct {
	ID        int64     `json:"-" db:"id"`
	UserID    string    `json:"user_id"`
	FriendID  string    `json:"friend_id"`
	Status    string    `json:"status"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

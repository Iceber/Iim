package model

import "time"

type Group struct {
	ID           int64     `json:"-" db:"id"`
	GroupID      string    `json:"group_id"`
	Name         string    `json:"name"`
	CreatorID    string    `json:"creator_id"`
	UserNum      string    `json:"user_num"`
	OwnerID      string    `json:"owner_id"`
	Introduction string    `json:"introduction"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type GroupUser struct {
	ID            int64     `json:"-" db:"id"`
	GroupID       string    `json:"group_id"`
	UserID        string    `json:"user_id"`
	GroupNickname string    `json:"group_nickname"`
	MessageAck    int64     `json:"message_ack"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

const (
	GroupRequestStatusPending = "applying"
	GroupStatusApproved       = "approved"
)

type GroupRequest struct {
	ID        int64     `json:"-" db:"id"`
	UserID    string    `json:"user_id"`
	GroupID   string    `json:"group_id"`
	Status    string    `json:"status"`
	ExpiredAt string    `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

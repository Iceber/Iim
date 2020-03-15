package model

import "time"

type User struct {
	ID        int64     `json:"-" db:"id"`
	UserID    string    `json:"user_id"`
	Nickname  string    `json:"nickname"`
	Sex       int32     `json:"sex"`
	AvatarURL string    `json:"avatar_url"`
	Extra     string    `json:"extra_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

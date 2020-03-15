package model

import "time"

type Device struct {
	ID        int64     `json:"-" db:"id"`
	DeviceID  int64     `json:"device_id"`
	UserID    string    `json:"user_id"`
	Type      int32     `json:"type"`
	Status    int32     `json:"status"`
	Extra     string    `json:"extra"`
	ConnAddr  string    `json:"-" db:"conn_addr"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

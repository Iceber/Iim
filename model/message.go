package model

import "time"

type Message struct {
	ID             int64  `json:"-" db:"id"`
	SenderID       string `json:"sender_id"`
	SenderDeviceID string `json:"sender_device_id"`
	ReceiverID     string `json:"receiver_id"`
	Type           int    `json:"type"`
	Content        string `json:"content"`

	Status    string    `json:"status"`
	Seq       int64     `json:"seq"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GroupMessage struct {
	ID      int64  `json:"-" db:"id"`
	GroupID string `json:"group_id"`

	SenderType     string `json:"sender_type"`
	SenderID       string `json:"sender_id"`
	SenderDeviceID string `json:"sender_device_id"`
	AlertedUserIDs string `json:"alterted_user_ids"`

	Type    string `json:"type"`
	Seq     int64  `json:"seq"`
	Content string `json:"content"`

	Status string `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

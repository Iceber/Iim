package service

import (
	"context"

	"github.com/iceber/iim/model"
)

type messageService struct{}

var MessageService = new(messageService)

func (ms *messageService) List(ctx context.Context, userID, friendID string, offset, limit int64) error {
}

func (ms *messageService) ListGroup(ctx context.Context, userID, groupID string) error {
}

func (ms *messageService) Send(ctx context.Context, message *model.Message) error {
}

func (ms *messageService) Ack(ctx context.Context, userID string, ack int64) error {
}

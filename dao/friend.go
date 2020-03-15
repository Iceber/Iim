package dao

import (
	"context"

	"github.com/iceber/iim/model"
)

func (d *dao) Friends(ctx context.Context, userID string) ([]model.Friend, error) {
}

func (d *dao) Friend(ctx context.Context, userID, friendID string) (*model.Friend, error) {
}

func (d *dao) Delete(ctx context.Context, userID, friendID string) error {
}

func (d *dao) Update(ctx context.Context, friendID string) error {
}

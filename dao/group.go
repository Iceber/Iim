package dao

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/iceber/iim/model"
	"github.com/jmoiron/sqlx"
)

type dao struct {
	DB *sqlx.DB
}

func (d *dao) CreateGroup(ctx context.Context, group *model.Group) error {
}

func (d *dao) Groups(ctx context.Context, userID string) ([]model.Group, error) {
	var total int
	query, args, err := sq.Select("COUNT(*)").From(model.TableGroup).Where(sq.Eq{"user_id": userID}).ToSql()
	if err != nil {
		return nil, err
	}
	if err := d.DB.GetContext(ctx, &total, query, args...); err != nil {
		return nil, err
	}

	groups := make([]model.Group, 0)
	query, args, err = sq.Select("*").From(model.TableGroup).Where(sq.Eq{"user_id": userID}).ToSql()
	if err != nil {
		return nil, err
	}
	if err := d.DB.SelectContext(ctx, &groups, query, args...); err != nil {
		return nil, err
	}
	return groups, nil
}

func (d *dao) Group(ctx context.Context, groupID string) (*model.Group, error) {
	query, args, err := sq.Select("*").From(model.TableGroup).Where(sq.Eq{"group_id": groupID}).ToSql()
	if err != nil {
		return nil, err
	}

	group := new(model.Group)
	if err := d.DB.GetContext(ctx, &group, query, args...); err != nil {
		return nil, err
	}
	return group, nil
}

func (d *dao) IsGroupUser(ctx context.Context, groupID, userID string) (bool, error) {
	query, args, err := sq.Select("COUNT(*)").From(model.TableGroup).Where(sq.Eq{"group_id": groupID, "user_id": userID, "is_active": true}).ToSql()
	if err != nil {
		return false, err
	}
	var count int
	if err := d.DB.GetContext(ctx, &count, query, args...); err != nil {
		return false, err
	}
	return count != 0, nil
}

func (d *dao) UpdateGroup(ctx context.Context, group *model.Group) error {
}

func (d *dao) DeleteGroup(ctx context.Context, groupID string) error {
}

func (d *dao) GroupAddUser(ctx context.Context, group *model.Group) error {
}

func (d *dao) GroupDelUser(ctx context.Context, group *model.Group) error {
}

package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ItemRepo interface {
	InsertItem(ctx context.Context, item *ItemDetail) (string, error)
	FindItem(ctx context.Context, id string) (*Item, error)
	FindItemBriefs(ctx context.Context, date *time.Time) ([]*Item, error)
	DeleteItemForUser(ctx context.Context, id string, uid string) error
}

type ItemUseCase struct {
	repo   ItemRepo
	logger *log.Helper
}

func NewItemUseCase(repo ItemRepo, logger log.Logger) *ItemUseCase {
	return &ItemUseCase{repo: repo, logger: log.NewHelper(logger)}
}

func (u *ItemUseCase) AddItem(ctx context.Context, item *ItemDetail) (err error) {
	_, err = u.repo.InsertItem(ctx, item)
	return
}

func (u *ItemUseCase) GetBriefs(ctx context.Context, date *time.Time) (items []*Item, err error) {
	return u.repo.FindItemBriefs(ctx, date)
}

func (u *ItemUseCase) GetDetail(ctx context.Context, id string) (item *Item, err error) {
	return u.repo.FindItem(ctx, id)
}

func (u *ItemUseCase) DeleteItem(ctx context.Context, id string, uid string) error {
	return u.repo.DeleteItemForUser(ctx, id, uid)
}

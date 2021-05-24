package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ItemRepo interface {
	InsertItem(context.Context, *ItemDetail) (string, error)
	FindItem(context.Context, string) (*Item, error)
	FindItemBriefs(context.Context, *time.Time) ([]*Item, error)
	DeleteItemWithUID(context.Context, string, string) error
}

type ItemUsecase struct {
	repo ItemRepo
	log  *log.Helper
}

func NewItemUsecase(repo ItemRepo, logger log.Logger) *ItemUsecase {
	return &ItemUsecase{repo: repo, log: log.NewHelper("usecase/item", logger)}
}

func (u *ItemUsecase) GetBriefs(ctx context.Context, date *time.Time) (items []*Item, err error) {
	return u.repo.FindItemBriefs(ctx, date)
}

func (u *ItemUsecase) GetDetail(ctx context.Context, id string) (item *Item, err error) {
	return u.repo.FindItem(ctx, id)
}

func (u *ItemUsecase) AddItem(ctx context.Context, item *ItemDetail) (err error) {
	_, err = u.repo.InsertItem(ctx, item)
	return
}

func (u *ItemUsecase) DeleteItem(ctx context.Context, id string, uid string) error {
	return u.repo.DeleteItemWithUID(ctx, id, uid)
}

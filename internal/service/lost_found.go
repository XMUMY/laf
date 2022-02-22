package service

import (
	"context"
	"time"

	auth "github.com/XMUMY/api/core/auth/v4"
	v4 "github.com/XMUMY/lost_found/api/lost_found/v4"
	"github.com/XMUMY/lost_found/internal/biz"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LostAndFoundService struct {
	v4.UnimplementedLostAndFoundServer
	authClient  *auth.Client
	itemUseCase *biz.ItemUseCase
}

func NewLostAndFoundService(authClient *auth.Client, itemUseCase *biz.ItemUseCase) *LostAndFoundService {
	return &LostAndFoundService{
		authClient:  authClient,
		itemUseCase: itemUseCase,
	}
}

func (s *LostAndFoundService) GetBriefs(ctx context.Context, req *v4.GetBriefsReq) (resp *v4.GetBriefsResp, err error) {
	var date time.Time
	if req.Before == nil {
		date = time.Now()
	} else {
		date = req.Before.AsTime()
	}

	items, err := s.itemUseCase.GetBriefs(ctx, &date)
	if err != nil {
		return
	}

	resp = &v4.GetBriefsResp{}
	for _, item := range items {
		resp.Briefs = append(resp.Briefs, &v4.LostAndFoundBrief{
			Id:       item.Id.Hex(),
			Uid:      item.Uid,
			Type:     item.Type,
			Name:     item.Name,
			Time:     timestamppb.New(item.Timestamp),
			Location: item.Location,
		})
	}

	return
}

func (s *LostAndFoundService) GetDetail(ctx context.Context, req *v4.GetDetailReq) (resp *v4.LostAndFoundDetail, err error) {
	item, err := s.itemUseCase.GetDetail(ctx, req.Id)
	if err != nil {
		return
	}

	resp = &v4.LostAndFoundDetail{
		Uid:         item.Uid,
		Type:        item.Type,
		Name:        item.Name,
		Time:        timestamppb.New(item.Timestamp),
		Location:    item.Location,
		Description: item.Description,
		Contacts:    item.Contacts,
	}

	return
}

func (s *LostAndFoundService) AddItem(ctx context.Context, req *v4.AddItemReq) (resp *empty.Empty, err error) {
	authed, err := s.authClient.TryAuthenticate(ctx)
	if err != nil {
		return
	}

	err = req.Validate()
	if err != nil {
		return
	}

	resp = &empty.Empty{}
	err = s.itemUseCase.AddItem(ctx, &biz.ItemDetail{
		ItemBrief: biz.ItemBrief{
			Uid:       authed.Uid,
			Type:      req.Type,
			Name:      req.Name,
			Timestamp: req.Time.AsTime(),
			Location:  req.Location,
		},
		Description: req.Description,
		Contacts:    req.Contacts,
	})

	return
}

func (s *LostAndFoundService) DeleteItem(ctx context.Context, req *v4.DeleteItemReq) (resp *empty.Empty, err error) {
	authed, err := s.authClient.TryAuthenticate(ctx)
	if err != nil {
		return
	}

	err = req.Validate()
	if err != nil {
		return
	}

	resp = &empty.Empty{}
	err = s.itemUseCase.DeleteItem(ctx, req.Id, authed.Uid)

	return
}

package handler

import (
	"context"
	"time"

	"github.com/XMUMY/api/core/auth"
	"github.com/XMUMY/lost_found/dao"
	"github.com/XMUMY/lost_found/model"
	"github.com/XMUMY/lost_found/proto/lost_found"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
)

type LostAndFound struct {
	dao *dao.Dao
}

func New() *LostAndFound {
	return &LostAndFound{
		dao: dao.New(),
	}
}

func (s *LostAndFound) GetBriefs(ctx context.Context, req *lostfound.GetBriefsReq, resp *lostfound.GetBriefsResp) error {
	var date time.Time
	if req.Before.Seconds == 0 {
		date = time.Now()
	} else {
		date, _ = ptypes.Timestamp(req.Before)
	}

	items, err := s.dao.FindItemBriefs(ctx, &date)
	if err != nil {
		return err
	}

	for _, item := range items {
		ts, _ := ptypes.TimestampProto(item.Timestamp)
		resp.Briefs = append(resp.Briefs, &lostfound.LostAndFoundBrief{
			Id:       item.Id.Hex(),
			Uid:      item.Uid,
			Type:     item.Type,
			Name:     item.Name,
			Time:     ts,
			Location: item.Location,
		})
	}

	return nil
}

func (s *LostAndFound) GetDetail(ctx context.Context, req *lostfound.GetDetailReq, resp *lostfound.LostAndFoundDetail) error {
	item, err := s.dao.FindItem(ctx, req.Id)
	if err != nil {
		return err
	}

	ts, _ := ptypes.TimestampProto(item.Timestamp)
	resp.Uid = item.Uid
	resp.Type = item.Type
	resp.Name = item.Name
	resp.Time = ts
	resp.Location = item.Location
	resp.Description = item.Description
	resp.Contacts = item.Contacts
	resp.Pictures = item.Pictures

	return nil
}

func (s *LostAndFound) AddItem(ctx context.Context, req *lostfound.AddItemReq, _ *empty.Empty) error {
	authed, _, err := auth.AuthenticateWithCampusIdPassword(ctx)
	if err != nil {
		return err
	}

	tm, _ := ptypes.Timestamp(req.Time)
	item := &model.LostAndFoundDetail{
		LostAndFoundBrief: model.LostAndFoundBrief{
			Uid:       authed.Uid,
			Type:      req.Type,
			Name:      req.Name,
			Timestamp: tm,
			Location:  req.Location,
		},
		Description: req.Description,
		Contacts:    req.Contacts,
		Pictures:    req.Pictures,
	}

	_, err = s.dao.InsertItem(ctx, item)
	return err
}

func (s *LostAndFound) DeleteItem(ctx context.Context, req *lostfound.DeleteItemReq, _ *empty.Empty) error {
	authed, _, err := auth.AuthenticateWithCampusIdPassword(ctx)
	if err != nil {
		return err
	}

	err = s.dao.DeleteItemWithUID(ctx, req.Id, authed.Uid)
	return err
}

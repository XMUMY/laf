package service

import (
	"context"
	"os"
	"time"

	auth "github.com/XMUMY/api/core/auth/v4"
	pb "github.com/XMUMY/lost_found/api/lost_found/v4"
	"github.com/XMUMY/lost_found/internal/biz"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LostAndFoundService struct {
	pb.UnimplementedLostAndFoundServer
	authClient  *auth.Client
	itemUsecase *biz.ItemUsecase
}

func NewLostAndFoundService(itemUsecase *biz.ItemUsecase) (svc *LostAndFoundService, err error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(os.Getenv("AUTH_ENDPOINT")),
		grpc.WithTimeout(2*time.Second),
	)

	svc = &LostAndFoundService{
		authClient:  auth.NewClient(conn),
		itemUsecase: itemUsecase,
	}
	return
}

func (s *LostAndFoundService) GetBriefs(ctx context.Context, req *pb.GetBriefsReq) (resp *pb.GetBriefsResp, err error) {
	var date time.Time
	if req.Before == nil {
		date = time.Now()
	} else {
		date = req.Before.AsTime()
	}

	items, err := s.itemUsecase.GetBriefs(ctx, &date)
	if err != nil {
		return
	}

	resp = &pb.GetBriefsResp{}
	for _, item := range items {
		resp.Briefs = append(resp.Briefs, &pb.LostAndFoundBrief{
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
func (s *LostAndFoundService) GetDetail(ctx context.Context, req *pb.GetDetailReq) (resp *pb.LostAndFoundDetail, err error) {
	item, err := s.itemUsecase.GetDetail(ctx, req.Id)
	if err != nil {
		return
	}

	resp = &pb.LostAndFoundDetail{
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
func (s *LostAndFoundService) AddItem(ctx context.Context, req *pb.AddItemReq) (resp *empty.Empty, err error) {
	authed, _, err := s.authClient.AuthenticateWithCampusIdPassword(ctx)
	if err != nil {
		return
	}

	err = req.Validate()
	if err != nil {
		return
	}

	resp = &empty.Empty{}
	err = s.itemUsecase.AddItem(ctx, &biz.ItemDetail{
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
func (s *LostAndFoundService) DeleteItem(ctx context.Context, req *pb.DeleteItemReq) (resp *empty.Empty, err error) {
	authed, _, err := s.authClient.AuthenticateWithCampusIdPassword(ctx)
	if err != nil {
		return
	}

	err = req.Validate()
	if err != nil {
		return
	}

	resp = &empty.Empty{}
	err = s.itemUsecase.DeleteItem(ctx, req.Id, authed.Uid)
	return
}

package model

import (
	"net/http"

	"github.com/XMUMY/lost_found/proto/lost_found"
	"github.com/micro/go-micro/v2/errors"
)

var (
	ItemNotFoundError = &errors.Error{
		Id:     lostfound.SvcID,
		Code:   http.StatusNotFound,
		Detail: "item not found",
		Status: http.StatusText(http.StatusNotFound),
	}

	InvalidItemIDError = errors.BadRequest(lostfound.SvcID, "invalid item ID")
)

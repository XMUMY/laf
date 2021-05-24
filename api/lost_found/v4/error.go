package v4

import (
	"github.com/go-kratos/kratos/v2/errors"
)

const domain = "api.xmux.xdea.io"

var (
	ItemNotFoundError = errors.NotFound(domain, ErrorReason_name[int32(ErrorReason_ITEM_NOT_FOUND)],
		"item not found")

	InvalidItemIDError = errors.BadRequest(domain, ErrorReason_name[int32(ErrorReason_INVALID_ITEM_ID)],
		"invalid item ID")
)

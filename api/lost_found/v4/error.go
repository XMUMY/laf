package v4

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ItemNotFoundError  *errors.Error
	InvalidItemIDError *errors.Error
)

func init() {
	file_api_lost_found_v4_error_reason_proto_init()
	ItemNotFoundError = ErrorItemNotFound("item not found")
	InvalidItemIDError = ErrorInvalidItemId("invalid item ID")
}

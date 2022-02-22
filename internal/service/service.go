package service

import (
	auth "github.com/XMUMY/api/core/auth/v4"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewLostAndFoundService, auth.NewDirectClient)

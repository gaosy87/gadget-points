package mock

import (
	"gadget-points/infrastructure/auth"
	"net/http"
)

// TokenInterface is a mock token interface
type TokenInterface struct {
	CreateTokenFn          func(userId uint64) (*auth.TokenDetails, error)
	ExtractTokenMetadataFn func(*http.Request) (*auth.AccessDetails, error)
}

func (f *TokenInterface) CreateToken(userid uint64) (*auth.TokenDetails, error) {
	return f.CreateTokenFn(userid)
}
func (f *TokenInterface) ExtractTokenMetadata(r *http.Request) (*auth.AccessDetails, error) {
	return f.ExtractTokenMetadataFn(r)
}

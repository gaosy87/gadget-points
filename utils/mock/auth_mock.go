package mock

import (
	"gadget-points/infrastructure/auth"
)

// AuthInterface is a mock auth interface
type AuthInterface struct {
	CreateAuthFn    func(uint64, *auth.TokenDetails) error
	FetchAuthFn     func(string) (uint64, error)
	DeleteRefreshFn func(string) error
	DeleteTokensFn  func(*auth.AccessDetails) error
}

func (f *AuthInterface) DeleteRefresh(refreshUuid string) error {
	return f.DeleteRefreshFn(refreshUuid)
}
func (f *AuthInterface) DeleteTokens(authD *auth.AccessDetails) error {
	return f.DeleteTokensFn(authD)
}
func (f *AuthInterface) FetchAuth(uuid string) (uint64, error) {
	return f.FetchAuthFn(uuid)
}
func (f *AuthInterface) CreateAuth(userId uint64, authD *auth.TokenDetails) error {
	return f.CreateAuthFn(userId, authD)
}

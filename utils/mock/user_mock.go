package mock

import (
	"gadget-points/domain/entity"
)

// UserAppInterface is a mock user app interface
type UserAppInterface struct {
	SaveUserFn                  func(*entity.User) (*entity.User, map[string]string)
	GetUsersFn                  func() ([]entity.User, error)
	GetUserFn                   func(uint64) (*entity.User, error)
	GetUserByEmailAndPasswordFn func(*entity.User) (*entity.User, map[string]string)
}

// SaveUser calls the SaveUserFn
func (u *UserAppInterface) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.SaveUserFn(user)
}

// GetUsersFn calls the GetUsers
func (u *UserAppInterface) GetUsers() ([]entity.User, error) {
	return u.GetUsersFn()
}

// GetUserFn calls the GetUser
func (u *UserAppInterface) GetUser(userId uint64) (*entity.User, error) {
	return u.GetUserFn(userId)
}

// GetUserByEmailAndPasswordFn calls the GetUserByEmailAndPassword
func (u *UserAppInterface) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.GetUserByEmailAndPasswordFn(user)
}

package application

import (
	"gadget-points/domain/entity"
	"gadget-points/domain/repository"
)

type orderApp struct {
	fr repository.OrderRepository
}

var _ OrderAppInterface = &orderApp{}

type OrderAppInterface interface {
	CreateOrder(order *entity.Order) error
}

func (f *orderApp) CreateOrder(order *entity.Order) error {
	return f.fr.CreateOrder(order)
}

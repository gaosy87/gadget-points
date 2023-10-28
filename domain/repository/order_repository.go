package repository

import "gadget-points/domain/entity"

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
}

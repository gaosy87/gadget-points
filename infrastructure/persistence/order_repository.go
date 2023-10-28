package persistence

import (
	"gadget-points/domain/entity"
	"gadget-points/domain/repository"
	"github.com/jinzhu/gorm"
	"strings"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db}
}

var _ repository.OrderRepository = &OrderRepo{}

func (r *OrderRepo) CreateOrder(order *entity.Order) error {
	err := r.db.Debug().Create(&order).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return err
		}

		return err
	}
	return nil
}

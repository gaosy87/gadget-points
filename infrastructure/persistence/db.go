package persistence

import (
	"fmt"
	"gadget-points/domain/entity"
	"gadget-points/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User     repository.UserRepository
	Product  repository.ProductRepository
	Agent    repository.AgentRepository
	Order    repository.OrderRepository
	Activity repository.ActivityRepository
	db       *gorm.DB
}

func NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open(dbDriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User:     NewUserRepository(db),
		Product:  NewProductRepository(db),
		Agent:    NewAgentRepository(db),
		Order:    NewOrderRepository(db),
		Activity: NewActivityRepository(db),
		db:       db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.Agent{}).Error
}

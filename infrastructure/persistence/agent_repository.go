package persistence

import (
	"errors"
	"gadget-points/domain/entity"
	"gadget-points/domain/repository"
	"github.com/jinzhu/gorm"
)

type AgentRepo struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) *AgentRepo {
	return &AgentRepo{db}
}

var _ repository.AgentRepository = &AgentRepo{}

func (r *AgentRepo) GetAgent(agentCode string) (*entity.Agent, error) {
	var agent entity.Agent
	err := r.db.Debug().Where("agent_code = ?", agentCode).Take(&agent).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("agent not found")
	}
	return &agent, nil
}

package repository

import "gadget-points/domain/entity"

type AgentRepository interface {
	GetAgent(agentCode string) (*entity.Agent, error)
}

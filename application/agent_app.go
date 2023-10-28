package application

import (
	"gadget-points/domain/entity"
	"gadget-points/domain/repository"
)

type agentApp struct {
	fr repository.AgentRepository
}

var _ AgentAppInterface = &agentApp{}

type AgentAppInterface interface {
	GetAgent(agentCode string) (*entity.Agent, error)
}

func (f *agentApp) GetAgent(agentCode string) (*entity.Agent, error) {
	return f.fr.GetAgent(agentCode)
}

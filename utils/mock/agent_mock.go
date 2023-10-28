package mock

import (
	"gadget-points/domain/entity"
)

// AgentAppInterface is a mock food app interface
type AgentAppInterface struct {
	GetAgentFn func(string) (*entity.Agent, error)
}

func (f *AgentAppInterface) GetAgent(agentCode string) (*entity.Agent, error) {
	return f.GetAgentFn(agentCode)
}

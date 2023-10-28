package application

import (
	"gadget-points/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type fakeAgentRepo struct{}

var (
	getAgentRepo func(string) (*entity.Agent, error)
)

func (f *fakeAgentRepo) GetAgent(agentCode string) (*entity.Agent, error) {
	return getAgentRepo(agentCode)
}

var agentAppFake AgentAppInterface = &fakeAgentRepo{}

func TestGetFood_Success(t *testing.T) {
	getAgentRepo = func(agentCode string) (*entity.Agent, error) {
		return &entity.Agent{
			ID:        1,
			AgentCode: "1001",
			AgentName: "xxx",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}
	agentCode := "1001"
	f, err := agentAppFake.GetAgent(agentCode)
	assert.Nil(t, err)
	assert.EqualValues(t, f.AgentCode, "1001")
	assert.EqualValues(t, f.AgentName, "xxx")
}

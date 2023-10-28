package interfaces

import (
	"gadget-points/application"
	"gadget-points/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Agent struct {
	agentApp application.AgentAppInterface
	tk       auth.TokenInterface
	rd       auth.AuthInterface
}

// NewAgent Agent constructor
func NewAgent(fApp application.AgentAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Agent {
	return &Agent{
		agentApp: fApp,
		rd:       rd,
		tk:       tk,
	}
}

func (t *Agent) GetAgent(c *gin.Context) {
	agentCode := c.Param("agent_code")
	if agentCode == "" {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	agent, err := t.agentApp.GetAgent(agentCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, agent)
}

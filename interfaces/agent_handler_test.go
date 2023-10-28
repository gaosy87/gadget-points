package interfaces

import (
	"encoding/json"
	"gadget-points/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAgent_Success(t *testing.T) {
	agentApp.GetAgentFn = func(string) (*entity.Agent, error) {
		return &entity.Agent{
			ID:        1,
			AgentCode: "1001",
			AgentName: "xxx",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	agentCode := "1001"
	req, err := http.NewRequest(http.MethodGet, "/agent/"+agentCode, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/agent/:agent_code", f.GetAgent)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var agent = make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &agent)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.Equal(t, rr.Code, 200)

	assert.EqualValues(t, agent["agent_name"], "xxx")
}

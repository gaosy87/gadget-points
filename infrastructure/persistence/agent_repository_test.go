package persistence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAgent_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	repo := NewAgentRepository(conn)

	f, saveErr := repo.GetAgent("1001")

	assert.Nil(t, saveErr)
	assert.EqualValues(t, f.AgentName, "xxx")
}

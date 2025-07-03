package api

import (
	"os"
	"testing"
	"time"

	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// new server for testing
func newTestServer(t *testing.T, store db.Store) *Server{
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}


func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	
	// Run tests
	os.Exit(m.Run())
}

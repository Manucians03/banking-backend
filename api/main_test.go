package api

import (
	"os"
	"testing"
	"time"

	db "github.com/Manucians03/banking-backend/db/sqlc"
	"github.com/Manucians03/banking-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenKey:      util.RandomString(32),
		TokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}

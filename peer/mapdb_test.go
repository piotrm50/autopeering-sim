package peer

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot initialize logger: %v", err)
	}
	logger = l.Sugar()
}

func TestMapDBPing(t *testing.T) {
	p := newTestPeer()
	db := NewMemoryDB(logger)

	time := time.Now()
	err := db.UpdateLastPing(p.ID(), p.Address(), time)
	require.NoError(t, err)

	assert.Equal(t, time.Unix(), db.LastPing(p.ID(), p.Address()).Unix())
}

func TestMapDBPong(t *testing.T) {
	p := newTestPeer()
	db := NewMemoryDB(logger)

	time := time.Now()
	err := db.UpdateLastPong(p.ID(), p.Address(), time)
	require.NoError(t, err)

	assert.Equal(t, time.Unix(), db.LastPong(p.ID(), p.Address()).Unix())
}

func TestMapDBPeer(t *testing.T) {
	p := newTestPeer()
	db := NewMemoryDB(logger)

	err := db.UpdatePeer(p)
	require.NoError(t, err)

	assert.Equal(t, p, db.Peer(p.ID()))
}

func TestMapDBRandomPeer(t *testing.T) {
	p := newTestPeer()
	db := NewMemoryDB(logger)

	require.NoError(t, db.UpdatePeer(p))
	require.NoError(t, db.UpdateLastPong(p.ID(), p.Address(), time.Now()))

	peers := db.GetRandomPeers(2, time.Second)
	assert.ElementsMatch(t, []*Peer{p}, peers)
}

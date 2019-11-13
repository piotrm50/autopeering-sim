package peer

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

// mapDB is a simple implementation of DB using a map.
type mapDB struct {
	mutex sync.RWMutex
	m     map[string]peerEntry

	log *zap.SugaredLogger

	wg      sync.WaitGroup
	closing chan struct{}
}

type peerEntry struct {
	data       []byte
	properties map[string]peerPropEntry
}

type peerPropEntry struct {
	lastPing, lastPong int64
}

// NewMemoryDB creates a new DB that uses a GO map.
func NewMemoryDB(log *zap.SugaredLogger) DB {
	db := &mapDB{
		m:       make(map[string]peerEntry),
		log:     log,
		closing: make(chan struct{}),
	}

	// start the expirer routine
	db.wg.Add(1)
	go db.expirer()

	return db
}

// Close closes the DB.
func (db *mapDB) Close() {
	db.log.Debugf("closing")

	close(db.closing)
	db.wg.Wait()
}

// LastPing returns that property for the given peer ID and address.
func (db *mapDB) LastPing(id ID, address string) time.Time {
	db.mutex.RLock()
	peerEntry := db.m[string(id.Bytes())]
	db.mutex.RUnlock()

	return time.Unix(peerEntry.properties[address].lastPing, 0)
}

// UpdateLastPing updates that property for the given peer ID and address.
func (db *mapDB) UpdateLastPing(id ID, address string, t time.Time) error {
	key := string(id.Bytes())

	db.mutex.Lock()
	peerEntry := db.m[key]
	if peerEntry.properties == nil {
		peerEntry.properties = make(map[string]peerPropEntry)
	}
	entry := peerEntry.properties[address]
	entry.lastPing = t.Unix()
	peerEntry.properties[address] = entry
	db.m[key] = peerEntry
	db.mutex.Unlock()

	return nil
}

// LastPong returns that property for the given peer ID and address.
func (db *mapDB) LastPong(id ID, address string) time.Time {
	db.mutex.RLock()
	peerEntry := db.m[string(id.Bytes())]
	db.mutex.RUnlock()

	return time.Unix(peerEntry.properties[address].lastPong, 0)
}

// UpdateLastPong updates that property for the given peer ID and address.
func (db *mapDB) UpdateLastPong(id ID, address string, t time.Time) error {
	key := string(id.Bytes())

	db.mutex.Lock()
	peerEntry := db.m[key]
	if peerEntry.properties == nil {
		peerEntry.properties = make(map[string]peerPropEntry)
	}
	entry := peerEntry.properties[address]
	entry.lastPong = t.Unix()
	peerEntry.properties[address] = entry
	db.m[key] = peerEntry
	db.mutex.Unlock()

	return nil
}

// UpdatePeer updates a peer in the database.
func (db *mapDB) UpdatePeer(p *Peer) error {
	data, err := p.Marshal()
	if err != nil {
		return err
	}
	key := string(p.ID().Bytes())

	db.mutex.Lock()
	peerEntry := db.m[key]
	peerEntry.data = data
	db.m[key] = peerEntry
	db.mutex.Unlock()

	return nil
}

// Peer retrieves a peer from the database.
func (db *mapDB) Peer(id ID) *Peer {
	db.mutex.RLock()
	peerEntry := db.m[string(id.Bytes())]
	db.mutex.RUnlock()

	return parsePeer(peerEntry.data)
}

// GetRandomPeers returns a random subset of n peers whose last ping is not too old.
func (db *mapDB) GetRandomPeers(n int, maxAge time.Duration) []*Peer {
	peers := make([]*Peer, 0)
	now := time.Now()

	db.mutex.RLock()
	for id, peerEntry := range db.m {
		p := parsePeer(peerEntry.data)
		if p == nil || id != string(p.ID().Bytes()) {
			continue
		}
		if now.Sub(db.LastPong(p.ID(), p.Address())) > maxAge {
			continue
		}

		peers = append(peers, p)
	}
	db.mutex.RUnlock()

	return randomSubset(peers, n)
}

func (db *mapDB) expirer() {
	defer db.wg.Done()

	// the expiring isn't triggert right away, to give the bootstrapping the chance to use older nodes
	tick := time.NewTicker(cleanupInterval)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			db.expirePeers()
		case <-db.closing:
			return
		}
	}
}

func (db *mapDB) expirePeers() {
	var (
		threshold = time.Now().Add(-peerExpiration).Unix()
		count     int
	)

	db.mutex.Lock()
	for id, peerEntry := range db.m {
		for address, peerPropEntry := range peerEntry.properties {
			if peerPropEntry.lastPong <= threshold {
				delete(peerEntry.properties, address)
			}
		}
		if len(peerEntry.properties) == 0 {
			delete(db.m, id)
			count++
		}
	}
	db.mutex.Unlock()

	db.log.Info("expired peers", "count", count)
}

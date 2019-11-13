package peer

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/iotaledger/goshimmer/packages/database"
)

const (
	// remove peers from DB, when the last received ping was older than this
	peerExpiration = 24 * time.Hour
	// interval in which expired peers are checked
	cleanupInterval = time.Hour
)

// DB is the peer database, storing previously seen peers and any collected
// properties of them.
type DB interface {
	// Peer retrieves a peer from the database.
	Peer(id ID) *Peer
	// UpdatePeer updates a peer in the database.
	UpdatePeer(p *Peer) error
	// GetRandomPeers returns a random subset of n peers whose last ping is not too old.
	GetRandomPeers(n int, maxAge time.Duration) []*Peer

	// LastPing returns that property for the given peer ID and address.
	LastPing(id ID, address string) time.Time
	// UpdateLastPing updates that property for the given peer ID and address.
	UpdateLastPing(id ID, address string, t time.Time) error

	// LastPong returns that property for the given peer ID and address.
	LastPong(id ID, address string) time.Time
	// UpdateLastPong updates that property for the given peer ID and address.
	UpdateLastPong(id ID, address string, t time.Time) error

	// Close closes the DB.
	Close()
}

type persistentDB struct {
	db database.Database
}

// Keys in the node database.
const (
	dbNodePrefix  = "n:"     // Identifier to prefix node entries with
	dbLocalPrefix = "local:" // Identifier to prefix local entries

	// These fields are stored per ID and address. Use nodeItemKey to create those keys.
	dbNodePing = "lastping"
	dbNodePong = "lastpong"

	// Local information is keyed by ID only. Use localItemKey to create those keys.
	dbLocalSeq = "seq"
)

func NewPersistentDB() DB {
	db, err := database.Get("peer")
	if err != nil {
		panic(err)
	}

	return &persistentDB{
		db: db,
	}
}

// Close closes the DB.
func (db *persistentDB) Close() {
	threshold := time.Now().Add(-peerExpiration).Unix()

	// remove TTL from all valid pong fields
	db.db.ForEachWithPrefix([]byte(dbNodePrefix), func(key []byte, value []byte) {
		if bytes.HasSuffix(key, []byte(dbNodePong)) {
			t := parseInt64(value)
			if t >= threshold {
				db.db.Set(key, value)
			}
		}
	})
}

// nodeKey returns the database key for a node record.
func nodeKey(id ID) []byte {
	return append([]byte(dbNodePrefix), id.Bytes()...)
}

func splitNodeKey(key []byte) (id ID, rest []byte) {
	if !bytes.HasPrefix(key, []byte(dbNodePrefix)) {
		return ID{}, nil
	}
	item := key[len(dbNodePrefix):]
	copy(id[:], item[:len(id)])
	return id, item[len(id)+1:]
}

// nodeItemKey returns the database key for a node metadata field.
func nodeItemKey(id ID, address string, field string) []byte {
	return bytes.Join([][]byte{nodeKey(id), []byte(address), []byte(field)}, []byte{':'})
}

func parseInt64(blob []byte) int64 {
	val, read := binary.Varint(blob)
	if read <= 0 {
		return 0
	}
	return val
}

// getInt64 retrieves an integer associated with a particular key.
func (db *persistentDB) getInt64(key []byte) int64 {
	blob, err := db.db.Get(key)
	if err != nil {
		return 0
	}
	return parseInt64(blob)
}

// setInt64 stores an integer in the given key.
func (db *persistentDB) setInt64(key []byte, n int64) error {
	blob := make([]byte, binary.MaxVarintLen64)
	blob = blob[:binary.PutVarint(blob, n)]
	return db.db.SetWithTTL(key, blob, peerExpiration)
}

// LastPing returns that property for the given peer ID and address.
func (db *persistentDB) LastPing(id ID, address string) time.Time {
	return time.Unix(db.getInt64(nodeItemKey(id, address, dbNodePing)), 0)
}

// UpdateLastPing updates that property for the given peer ID and address.
func (db *persistentDB) UpdateLastPing(id ID, address string, t time.Time) error {
	return db.setInt64(nodeItemKey(id, address, dbNodePing), t.Unix())
}

// LastPing returns that property for the given peer ID and address.
func (db *persistentDB) LastPong(id ID, address string) time.Time {
	return time.Unix(db.getInt64(nodeItemKey(id, address, dbNodePong)), 0)
}

// UpdateLastPing updates that property for the given peer ID and address.
func (db *persistentDB) UpdateLastPong(id ID, address string, t time.Time) error {
	return db.setInt64(nodeItemKey(id, address, dbNodePong), t.Unix())
}

func (db *persistentDB) UpdatePeer(p *Peer) error {
	data, err := p.Marshal()
	if err != nil {
		return err
	}
	return db.db.SetWithTTL(nodeKey(p.ID()), data, peerExpiration)
}

func parsePeer(data []byte) *Peer {
	p, err := Unmarshal(data)
	if err != nil {
		return nil
	}
	return p
}

func (db *persistentDB) Peer(id ID) *Peer {
	data, err := db.db.Get(nodeKey(id))
	if err != nil {
		return nil
	}
	return parsePeer(data)
}

func randomSubset(peers []*Peer, m int) []*Peer {
	if len(peers) <= m {
		return peers
	}

	result := make([]*Peer, 0, m)
	for i, p := range peers {
		if rand.Intn(len(peers)-i) < m-len(result) {
			result = append(result, p)
		}
	}
	return result
}

// GetRandomPeers retrieves random nodes to be used as potential bootstrap peers.
func (db *persistentDB) GetRandomPeers(n int, maxAge time.Duration) []*Peer {
	peers := make([]*Peer, 0)
	now := time.Now()

	err := db.db.ForEachWithPrefix([]byte(dbNodePrefix), func(key []byte, value []byte) {
		id, rest := splitNodeKey(key)
		if len(rest) > 0 {
			return
		}

		p := parsePeer(value)
		if p == nil || p.ID() != id {
			return
		}
		if now.Sub(db.LastPong(p.ID(), p.Address())) > maxAge {
			return
		}

		peers = append(peers, p)
	})
	if err != nil {
		return []*Peer{}
	}

	return randomSubset(peers, n)
}

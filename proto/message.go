package proto

import "github.com/golang/protobuf/proto"

// An enum for the different message types.
const (
	MPing MType = iota
	MPong
	MPeersRequest
	MPeersResponse
)

// MType is the type of message type enum.
type MType int

// Message extends the proto.Message interface with additional util functions.
type Message interface {
	proto.Message

	// Name returns the name of the corresponding message type for debugging.
	Name() string
	// Type returns the type of the corresponding message as an enum.
	Type() MType

	// Wrapper returns the corresponding wrapped message.
	Wrapper() *MessageWrapper
}

func (m *Ping) Name() string { return "PING" }
func (m *Ping) Type() MType  { return MPing }
func (m *Ping) Wrapper() *MessageWrapper {
	return &MessageWrapper{Message: &MessageWrapper_Ping{Ping: m}}
}

func (m *Pong) Name() string { return "PONG" }
func (m *Pong) Type() MType  { return MPong }
func (m *Pong) Wrapper() *MessageWrapper {
	return &MessageWrapper{Message: &MessageWrapper_Pong{Pong: m}}
}

func (m *PeersRequest) Name() string { return "PEERS_REQUEST" }
func (m *PeersRequest) Type() MType  { return MPeersRequest }
func (m *PeersRequest) Wrapper() *MessageWrapper {
	return &MessageWrapper{Message: &MessageWrapper_PeersRequest{PeersRequest: m}}
}

func (m *PeersResponse) Name() string { return "PEERS_RESPONSE" }
func (m *PeersResponse) Type() MType  { return MPeersResponse }
func (m *PeersResponse) Wrapper() *MessageWrapper {
	return &MessageWrapper{Message: &MessageWrapper_PeersResponse{PeersResponse: m}}
}

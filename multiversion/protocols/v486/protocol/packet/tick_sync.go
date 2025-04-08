package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// TickSync is sent by the client and the server to maintain a synchronized, server-authoritative tick between
// the client and the server. The client sends this packet first, and the server should reply with another one
// of these packets, including the response time.
type TickSync struct {
	ClientRequestTimestamp   int64
	ServerReceptionTimestamp int64
}

// ID ...
func (*TickSync) ID() uint32 {
	return IDTickSync
}

// Marshal ...
func (pk *TickSync) Marshal(io protocol.IO) {
	io.Int64(&pk.ClientRequestTimestamp)
	io.Int64(&pk.ServerReceptionTimestamp)
}

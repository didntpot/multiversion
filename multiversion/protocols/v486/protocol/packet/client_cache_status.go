package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ClientCacheStatus is sent by the client to the server at the start of the game. It is sent to let the
// server know if it supports the client-side blob cache. Clients such as Nintendo Switch do not support the
// cache, and attempting to use it anyway will fail.
type ClientCacheStatus struct {
	Enabled bool
}

// ID ...
func (pk *ClientCacheStatus) ID() uint32 {
	return IDClientCacheStatus
}

// Marshal ...
func (pk *ClientCacheStatus) Marshal(io protocol.IO) {
	io.Bool(&pk.Enabled)
}

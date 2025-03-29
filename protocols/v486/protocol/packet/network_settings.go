package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// NetworkSettings is sent by the server to update a variety of network settings. These settings modify the
// way packets are sent over the network stack.
type NetworkSettings struct {
	CompressionThreshold uint16
}

// ID ...
func (*NetworkSettings) ID() uint32 {
	return IDNetworkSettings
}

// Marshal ...
func (pk *NetworkSettings) Marshal(io protocol.IO) {
	io.Uint16(&pk.CompressionThreshold)
}

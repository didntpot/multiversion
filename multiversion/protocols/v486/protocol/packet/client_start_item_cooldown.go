package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ClientStartItemCooldown is sent by the client to the server to initiate a cooldown on an item. The purpose of this
// packet isn't entirely clear.
type ClientStartItemCooldown struct {
	Category string
	Duration int32
}

// ID ...
func (*ClientStartItemCooldown) ID() uint32 {
	return IDClientStartItemCooldown
}

// Marshal ...
func (pk *ClientStartItemCooldown) Marshal(io protocol.IO) {
	io.String(&pk.Category)
	io.Varint32(&pk.Duration)
}

package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ItemComponent is sent by the server to attach client-side components to a custom item.
type ItemComponent struct {
	Items []legacyprotocol.ItemComponentEntry
}

// ID ...
func (*ItemComponent) ID() uint32 {
	return IDItemComponent
}

// Marshal ...
func (pk *ItemComponent) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Items)
}

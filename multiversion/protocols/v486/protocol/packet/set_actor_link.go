package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetActorLink is sent by the server to initiate an entity link client-side, meaning one entity will start
// riding another.
type SetActorLink struct {
	EntityLink legacyprotocol.EntityLink
}

// ID ...
func (*SetActorLink) ID() uint32 {
	return IDSetActorLink
}

// Marshal ...
func (pk *SetActorLink) Marshal(io protocol.IO) {
	protocol.Single(io, &pk.EntityLink)
}

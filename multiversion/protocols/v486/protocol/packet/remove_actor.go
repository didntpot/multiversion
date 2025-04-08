package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// RemoveActor is sent by the server to remove an entity that currently exists in the world from the client-
// side. Sending this packet if the client cannot already see this entity will have no effect.
type RemoveActor struct {
	EntityUniqueID int64
}

// ID ...
func (*RemoveActor) ID() uint32 {
	return IDRemoveActor
}

// Marshal ...
func (pk *RemoveActor) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
}

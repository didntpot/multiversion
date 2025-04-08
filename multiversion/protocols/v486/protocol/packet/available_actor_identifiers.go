package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AvailableActorIdentifiers is sent by the server at the start of the game to let the client know all
// entities that are available on the server.
type AvailableActorIdentifiers struct {
	SerialisedEntityIdentifiers []byte
}

// ID ...
func (*AvailableActorIdentifiers) ID() uint32 {
	return IDAvailableActorIdentifiers
}

// Marshal ...
func (pk *AvailableActorIdentifiers) Marshal(io protocol.IO) {
	io.Bytes(&pk.SerialisedEntityIdentifiers)
}

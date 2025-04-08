package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UpdateAttributes is sent by the server to update an amount of attributes of any entity in the world. These
// attributes include ones such as the health or the movement speed of the entity.
type UpdateAttributes struct {
	EntityRuntimeID uint64
	Attributes      []legacyprotocol.Attribute
	Tick            uint64
}

// ID ...
func (*UpdateAttributes) ID() uint32 {
	return IDUpdateAttributes
}

// Marshal ...
func (pk *UpdateAttributes) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	protocol.Slice(io, &pk.Attributes)
	io.Varuint64(&pk.Tick)
}

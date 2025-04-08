package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetLastHurtBy is sent by the server to let the client know what entity type it was last hurt by. At this
// moment, the packet is useless and should not be used. There is no behaviour that depends on if this
// packet is sent or not.
type SetLastHurtBy struct {
	EntityType int32
}

// ID ...
func (*SetLastHurtBy) ID() uint32 {
	return IDSetLastHurtBy
}

// Marshal ...
func (pk *SetLastHurtBy) Marshal(io protocol.IO) {
	io.Varint32(&pk.EntityType)
}

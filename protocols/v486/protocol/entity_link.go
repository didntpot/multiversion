package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// EntityLink is a link between two entities, typically being one entity riding another.
type EntityLink struct {
	protocol.EntityLink
}

// Marshal ...
func (x *EntityLink) Marshal(io protocol.IO) {
	io.Varint64(&x.RiddenEntityUniqueID)
	io.Varint64(&x.RiderEntityUniqueID)
	io.Uint8(&x.Type)
	io.Bool(&x.Immediate)
	io.Bool(&x.RiderInitiated)
}

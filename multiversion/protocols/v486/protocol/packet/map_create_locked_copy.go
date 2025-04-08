package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MapCreateLockedCopy is sent by the server to create a locked copy of one map into another map. In vanilla,
// it is used in the cartography table to create a map that is locked and cannot be modified.
type MapCreateLockedCopy struct {
	OriginalMapID int64
	NewMapID      int64
}

// ID ...
func (*MapCreateLockedCopy) ID() uint32 {
	return IDMapCreateLockedCopy
}

// Marshal ...
func (pk *MapCreateLockedCopy) Marshal(io protocol.IO) {
	io.Varint64(&pk.OriginalMapID)
	io.Varint64(&pk.NewMapID)
}

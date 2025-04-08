package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BlockToEntityTransition = iota + 1
	EntityToBlockTransition
)

// UpdateBlockSynced is sent by the server to synchronise the falling of a falling block entity with the
// transitioning back and forth from and to a solid block. It is used to prevent the entity from flickering,
// and is used in places such as the pushing of blocks with pistons.
type UpdateBlockSynced struct {
	Position          protocol.BlockPos
	NewBlockRuntimeID uint32
	Flags             uint32
	Layer             uint32
	EntityUniqueID    int64
	TransitionType    uint64
}

// ID ...
func (*UpdateBlockSynced) ID() uint32 {
	return IDUpdateBlockSynced
}

// Marshal ...
func (pk *UpdateBlockSynced) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	io.Varuint32(&pk.NewBlockRuntimeID)
	io.Varuint32(&pk.Flags)
	io.Varuint32(&pk.Layer)
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.TransitionType)
}

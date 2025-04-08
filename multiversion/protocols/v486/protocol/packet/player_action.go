package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerAction is sent by the client when it executes any action, for example starting to sprint, swim,
// starting the breaking of a block, dropping an item, etc.
type PlayerAction struct {
	EntityRuntimeID uint64
	ActionType      int32
	BlockPosition   protocol.BlockPos
	BlockFace       int32
}

// ID ...
func (*PlayerAction) ID() uint32 {
	return IDPlayerAction
}

// Marshal ...
func (pk *PlayerAction) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.ActionType)
	io.UBlockPos(&pk.BlockPosition)
	io.Varint32(&pk.BlockFace)
}

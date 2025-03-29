package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BlockUpdateNeighbours = 1 << iota
	BlockUpdateNetwork
	BlockUpdateNoGraphics
	BlockUpdatePriority
)

// UpdateBlock is sent by the server to update a block client-side, without resending the entire chunk that
// the block is located in. It is particularly useful for small modifications like block breaking/placing.
type UpdateBlock struct {
	Position          protocol.BlockPos
	NewBlockRuntimeID uint32
	Flags             uint32
	Layer             uint32
}

// ID ...
func (*UpdateBlock) ID() uint32 {
	return IDUpdateBlock
}

// Marshal ...
func (pk *UpdateBlock) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	io.Varuint32(&pk.NewBlockRuntimeID)
	io.Varuint32(&pk.Flags)
	io.Varuint32(&pk.Layer)
}

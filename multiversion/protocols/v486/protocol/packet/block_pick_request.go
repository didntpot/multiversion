package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// BlockPickRequest is sent by the client when it requests to pick a block in the world and place its item in
// their inventory.
type BlockPickRequest struct {
	Position    protocol.BlockPos
	AddBlockNBT bool
	HotBarSlot  byte
}

// ID ...
func (*BlockPickRequest) ID() uint32 {
	return IDBlockPickRequest
}

// Marshal ...
func (pk *BlockPickRequest) Marshal(io protocol.IO) {
	io.BlockPos(&pk.Position)
	io.Bool(&pk.AddBlockNBT)
	io.Uint8(&pk.HotBarSlot)
}

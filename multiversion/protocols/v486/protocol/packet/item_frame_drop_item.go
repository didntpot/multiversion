package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ItemFrameDropItem is sent by the client when it takes an item out of an item frame.
type ItemFrameDropItem struct {
	Position protocol.BlockPos
}

// ID ...
func (*ItemFrameDropItem) ID() uint32 {
	return IDItemFrameDropItem
}

// Marshal ...
func (pk *ItemFrameDropItem) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
}

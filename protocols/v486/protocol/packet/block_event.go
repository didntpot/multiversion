package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BlockEventChangeChestState = 1
)

// BlockEvent is sent by the server to initiate a certain event that has something to do with blocks in
// specific, for example opening a chest.
type BlockEvent struct {
	Position  protocol.BlockPos
	EventType int32
	EventData int32
}

// ID ...
func (*BlockEvent) ID() uint32 {
	return IDBlockEvent
}

// Marshal ...
func (pk *BlockEvent) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	io.Varint32(&pk.EventType)
	io.Varint32(&pk.EventData)
}

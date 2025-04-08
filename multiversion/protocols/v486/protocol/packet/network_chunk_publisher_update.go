package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// NetworkChunkPublisherUpdate is sent by the server to change the point around which chunks are and remain
// loaded. This is useful for mini-game servers, where only one area is ever loaded, in which case the
// NetworkChunkPublisherUpdate packet can be sent in the middle of it, so that no chunks ever need to be
// additionally sent during the course of the game.
// In reality, the packet is not extraordinarily useful, and most servers just send it constantly at the
// position of the player.
// If the packet is not sent at all, no chunks will be shown to the player, regardless of where they are sent.
type NetworkChunkPublisherUpdate struct {
	Position protocol.BlockPos
	Radius   uint32
}

// ID ...
func (*NetworkChunkPublisherUpdate) ID() uint32 {
	return IDNetworkChunkPublisherUpdate
}

// Marshal ...
func (pk *NetworkChunkPublisherUpdate) Marshal(io protocol.IO) {
	io.BlockPos(&pk.Position)
	io.Varuint32(&pk.Radius)
}

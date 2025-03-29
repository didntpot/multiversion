package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ContainerOpen is sent by the server to open a container client-side. This container must be physically
// present in the world, for the packet to have any effect. Unlike Java Edition, Bedrock Edition requires that
// chests for example must be present and in range to open its inventory.
type ContainerOpen struct {
	WindowID                byte
	ContainerType           byte
	ContainerPosition       protocol.BlockPos
	ContainerEntityUniqueID int64
}

// ID ...
func (*ContainerOpen) ID() uint32 {
	return IDContainerOpen
}

// Marshal ...
func (pk *ContainerOpen) Marshal(io protocol.IO) {
	io.Uint8(&pk.WindowID)
	io.Uint8(&pk.ContainerType)
	io.UBlockPos(&pk.ContainerPosition)
	io.Varint64(&pk.ContainerEntityUniqueID)
}

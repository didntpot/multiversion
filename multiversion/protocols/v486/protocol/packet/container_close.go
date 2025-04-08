package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ContainerClose is sent by the server to close a container the player currently has opened, which was opened
// using the ContainerOpen packet, or by the client to tell the server it closed a particular container, such
// as the crafting grid.
type ContainerClose struct {
	WindowID   byte
	ServerSide bool
}

// ID ...
func (*ContainerClose) ID() uint32 {
	return IDContainerClose
}

// Marshal ...
func (pk *ContainerClose) Marshal(io protocol.IO) {
	io.Uint8(&pk.WindowID)
	io.Bool(&pk.ServerSide)
}

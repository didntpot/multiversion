package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CommandRequest is sent by the client to request the execution of a server-side command. Although some
// servers support sending commands using the Text packet, this packet is guaranteed to have the correct
// result.
type CommandRequest struct {
	CommandLine   string
	CommandOrigin protocol.CommandOrigin
	Internal      bool
}

// ID ...
func (*CommandRequest) ID() uint32 {
	return IDCommandRequest
}

// Marshal ...
func (pk *CommandRequest) Marshal(io protocol.IO) {
	io.String(&pk.CommandLine)
	protocol.CommandOriginData(io, &pk.CommandOrigin)
	io.Bool(&pk.Internal)
}

package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetCommandsEnabled is sent by the server to enable or disable the ability to execute commands for the
// client. If disabled, the client itself will stop the execution of commands.
type SetCommandsEnabled struct {
	Enabled bool
}

// ID ...
func (*SetCommandsEnabled) ID() uint32 {
	return IDSetCommandsEnabled
}

// Marshal ...
func (pk *SetCommandsEnabled) Marshal(io protocol.IO) {
	io.Bool(&pk.Enabled)
}

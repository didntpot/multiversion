package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SettingsCommand is sent by the client when it changes a setting in the settings that results in the issuing
// of a command to the server, such as when Show Coordinates is enabled.
type SettingsCommand struct {
	CommandLine    string
	SuppressOutput bool
}

// ID ...
func (*SettingsCommand) ID() uint32 {
	return IDSettingsCommand
}

// Marshal ...
func (pk *SettingsCommand) Marshal(io protocol.IO) {
	io.String(&pk.CommandLine)
	io.Bool(&pk.SuppressOutput)
}

package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetDefaultGameType is sent by the client when it toggles the default game type in the settings UI, and is
// sent by the server when it actually changes the default game type, resulting in the toggle being changed
// in the settings UI.
type SetDefaultGameType struct {
	GameType int32
}

// ID ...
func (*SetDefaultGameType) ID() uint32 {
	return IDSetDefaultGameType
}

// Marshal ...
func (pk *SetDefaultGameType) Marshal(io protocol.IO) {
	io.Varint32(&pk.GameType)
}

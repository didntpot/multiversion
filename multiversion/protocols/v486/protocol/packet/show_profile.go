package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ShowProfile is sent by the server to show the XBOX Live profile of one player to another.
type ShowProfile struct {
	XUID string
}

// ID ...
func (*ShowProfile) ID() uint32 {
	return IDShowProfile
}

// Marshal ...
func (pk *ShowProfile) Marshal(io protocol.IO) {
	io.String(&pk.XUID)
}

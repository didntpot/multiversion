package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Disconnect may be sent by the server to disconnect the client using an optional message to send as the
// disconnect screen.
type Disconnect struct {
	HideDisconnectionScreen bool
	Message                 string
}

// ID ...
func (*Disconnect) ID() uint32 {
	return IDDisconnect
}

// Marshal ...
func (pk *Disconnect) Marshal(io protocol.IO) {
	io.Bool(&pk.HideDisconnectionScreen)
	if !pk.HideDisconnectionScreen {
		io.String(&pk.Message)
	}
}

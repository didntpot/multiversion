package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Login is sent when the client initially tries to join the server. It is the first packet sent and contains
// information specific to the player.
type Login struct {
	ClientProtocol    int32
	ConnectionRequest []byte
}

// ID ...
func (*Login) ID() uint32 {
	return IDLogin
}

// Marshal ...
func (pk *Login) Marshal(io protocol.IO) {
	io.BEInt32(&pk.ClientProtocol)
	io.ByteSlice(&pk.ConnectionRequest)
}
